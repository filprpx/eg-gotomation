# Netbox Client Refactor Plan

## Goal

Restructure the Netbox client in `internal/netbox` to support the full Netbox API surface (50+ resources) without creating an unmaintainable single package, while preserving the ergonomic `client.Device.List(ctx)` syntax and establishing reusable patterns for future system clients (ServiceNow, etc.).

---

## Current State

- All Netbox code lives in `package netbox` inside `internal/netbox/`.
- The `NetboxClient` embeds grouped API structs (`DCIM`, `IPAM`) which in turn embed resource structs (`Device`, `Cable`, etc.).
- This creates clean call sites (`client.DCIM.Device.List`) but forces everything into one package due to Go's prohibition on import cycles.
- At ~10 files and partial coverage, the package already feels cluttered. Scaling to full API coverage would make it unmanageable.

---

## Target Architecture

### 1. Shared Transport Layer: `internal/transport`

A cross-system package that defines the common HTTP contract and a reusable base implementation.

```
internal/transport/
  doer.go       # Doer interface
  client.go     # BaseClient concrete implementation
  errors.go     # Shared response helpers and error formatting
```

**`Doer` interface** — the bare-bones contract every system client satisfies:

```go
type Doer interface {
    Get(ctx context.Context, path string) (*http.Response, error)
    Post(ctx context.Context, path string, body io.Reader) (*http.Response, error)
    Patch(ctx context.Context, path string, body io.Reader) (*http.Response, error)
    Delete(ctx context.Context, path string) (*http.Response, error)
}
```

**`BaseClient`** handles the repetitive plumbing:
- Base URL resolution
- Request timeout
- Default header injection
- Generic status-code error formatting

System-specific clients (Netbox, ServiceNow) embed `*BaseClient` and only configure their own auth.

### 2. System-Specific Packages

Each system gets its own tree. Resource logic lives in sub-packages that depend only on `transport.Doer`.

```
internal/netbox/
  client.go          # NetboxClient: embeds *transport.BaseClient, wires resource APIs
  auth.go            # Netbox-specific auth (token from env, inject into base headers)

  dcim/
    device.go        # Device, DeviceWrite, DeviceAPI
    cable.go         # Cable, CableAPI
    site.go          # Site, SiteAPI
    device_type.go   # DeviceType, DeviceTypesAPI
    device_role.go   # DeviceRole, DeviceRolesAPI
    rack.go          # ... etc
    ...

  ipam/
    prefix.go
    ipaddress.go
    vrf.go
    ...

  circuits/
    ...

  tenancy/
    ...

  virtualization/
    ...

internal/servicenow/
  client.go          # Also embeds *transport.BaseClient
  auth.go            # ServiceNow-specific auth (basic auth, OAuth, etc.)
  incident/
    ...
```

### 3. Flat API Surface

Drop the `DCIM` / `IPAM` grouping. Resource names are distinctive enough.

```go
client, _ := netbox.NewClient()
client.Device.List(ctx)
client.Site.Get(ctx, 42)
client.Prefix.Create(ctx, &prefix)
```

### 4. Why This Avoids Import Cycles

- `transport` → depends on nothing
- `netbox/dcim` → depends on `transport`
- `netbox` (root) → depends on `transport` + `netbox/dcim` + `netbox/ipam`
- No package imports its parent. Clean DAG.

---

## Tradeoffs and Decisions

### Decision: Include `BaseClient` in `transport`, not just the interface

**Rationale:** Netbox and ServiceNow clients become tiny (~20 lines). All HTTP boilerplate lives in one place. The interface alone would force copy-pasting the same `http.NewRequestWithContext`, header merging, and error formatting in every new system. If an exotic client (gRPC, heavy middleware) appears later, we can introduce a second interface — the cost of that future flexibility today is too high.

### Decision: Auth plugs into `BaseClient` at construction time

**Rationale:** No separate `Auth()` method. The Netbox `NewClient()` loads the API key from env and injects it into `BaseClient.Header`. ServiceNow will do the same with its own mechanism (basic auth token, OAuth bearer, etc.). Auth becomes a configuration detail, not a runtime state mutation.

### Decision: Generic response helpers live in `transport`, list response types stay in system packages

**Rationale:** `IsSuccess`, `IsError`, and body-reading error formatting are 100% generic. `ApiListResponse[T]` is Netbox-specific (`count`, `next`, `previous`, `results`). ServiceNow may paginate differently, so list wrappers stay local to the system package unless proven reusable.

### Decision: Package name is `internal/transport`

**Rationale:** `transport` signals "this is how we move bytes, not business logic." Alternatives considered: `client` (too generic, clashes with `netbox.Client`), `api` (implies business API types), `httpkit` (accurate but verbose). `transport` is concise and intent-revealing.

---

## Implementation Steps

### Phase 1: Bootstrap `internal/transport`

1. Create `internal/transport/doer.go` with the `Doer` interface.
2. Create `internal/transport/client.go` with `BaseClient`:
   - Fields: `BaseURL string`, `HTTP *http.Client`, `Header http.Header`
   - Methods implementing `Doer`
3. Create `internal/transport/errors.go` with:
   - `IsSuccess(res *http.Response) bool`
   - `IsError(res *http.Response) bool`
   - `APIError(res *http.Response) error` — body-reading, HTML guard, truncation logic extracted from current `NetboxClient.apiError()`

### Phase 2: Restructure `internal/netbox`

1. Create `internal/netbox/dcim/`.
   - Move `Device`, `DeviceWrite`, `DeviceAPI` → `dcim/device.go`.
   - Change `DeviceAPI.client` field type from `*NetboxClient` to `transport.Doer`.
   - Move `Cable`, `CableAPI` → `dcim/cable.go`.
   - Move `Site`, `SiteAPI` → `dcim/site.go`.
   - Move `DeviceType`, `DeviceTypesAPI` → `dcim/device_type.go`.
   - Move `DeviceRole`, `DeviceRolesAPI` → `dcim/device_role.go`.
   - Update constructors: `NewDeviceAPI(d transport.Doer) *DeviceAPI`, etc.
   - Use `transport.IsError` and `transport.APIError` instead of current helpers.
2. Create `internal/netbox/ipam/`.
   - Move IPAM types and future resources.
3. Rewrite `internal/netbox/client.go`.
   - `NetboxClient` embeds `*transport.BaseClient`.
   - Remove `get`, `post`, `patch`, `delete` methods (now handled by `BaseClient`).
   - Flatten API surface: `Device`, `Cable`, `Site`, `DeviceType`, `DeviceRole`, `Prefix`, etc. directly on `NetboxClient`.
   - `NewClient()` loads env vars, creates `BaseClient`, then constructs each resource API with `dcim.NewDeviceAPI(c)`, etc.
4. Rewrite `internal/netbox/auth.go`.
   - Load `NETBOX_API_KEY` and `NETBOX_URL`.
   - Configure `BaseClient` base URL and `Authorization` header.

### Phase 3: Update Callers

1. Update `cmd/rundeck/netbox/clone/main.go` and any other `cmd/` files.
   - Change `client.DCIM.Device.List(ctx)` → `client.Device.List(ctx)`.
   - Update imports if needed.

### Phase 4: Scale to Full API

1. For each new Netbox resource:
   - Create a file in the appropriate sub-package (`dcim/`, `ipam/`, etc.).
   - Define `Read` struct, `Write` struct (if needed), and `*API` struct.
   - Implement CRUD methods using `transport.Doer`.
   - Wire constructor in `NetboxClient`.

### Phase 5 (Optional): Generic CRUD Helper

If the repetition of `List` / `Get` / `Create` / `Update` / `Delete` across 50+ resources becomes tedious:

```go
// In api/ or a helper inside transport
type ResourceAPI[T any] struct {
    doer   Doer
    path   string
}

func (a *ResourceAPI[T]) List(ctx context.Context) ([]T, error)   { ... }
func (a *ResourceAPI[T]) Get(ctx context.Context, id int) (*T, error)   { ... }
// etc.
```

Each sub-package would then only define its struct and a thin wrapper:

```go
// dcim/device.go
type DeviceAPI struct {
    *transport.ResourceAPI[Device]
}

func NewDeviceAPI(d transport.Doer) *DeviceAPI {
    return &DeviceAPI{transport.NewResourceAPI[Device](d, "/api/dcim/devices/")}
}
```

Evaluate after Phase 4 if the boilerplate justifies the abstraction.

---

## Future Systems

When adding ServiceNow (or any other system):

1. Create `internal/<system>/`.
2. Embed `*transport.BaseClient` in `<system>.Client`.
3. Write `<system>/auth.go` for its specific auth mechanism.
4. Create sub-packages for its resource domains.
5. Resource APIs depend only on `transport.Doer` — zero HTTP code duplication.

---

## Checklist

- [ ] Phase 1: `internal/transport` created and compiles.
- [ ] Phase 2: `internal/netbox` restructured into sub-packages.
- [ ] Phase 3: All `cmd/` callers updated and working.
- [ ] Phase 4: Existing resources (Device, Cable, Site, DeviceType, DeviceRole) ported.
- [ ] Phase 5 (optional): Generic `ResourceAPI[T]` evaluated and optionally implemented.
