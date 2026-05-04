# Netbox Client Refactor Plan (v2)

## Goal

Restructure the Netbox client in `internal/netbox` to support the full Netbox API surface (50+ resources) without creating an unmaintainable single package, while preserving the ergonomic `client.Device.List(ctx)` syntax, avoiding circular dependencies via a shared models package, and establishing reusable patterns for future system clients (ServiceNow, etc.).

---

## Current State

- All Netbox code lives in `package netbox` inside `internal/netbox/`.
- The `NetboxClient` embeds grouped API structs (`DCIM`, `IPAM`) which in turn embed resource structs (`Device`, `Cable`, etc.).
- This creates clean call sites (`client.DCIM.Device.List`) but forces everything into one package due to Go's prohibition on import cycles.
- At ~10 files and partial coverage, the package already feels cluttered. Scaling to full API coverage would make it unmanageable.
- Cross-domain references (e.g. `Device` → `IPAddress`, `Prefix` → `Site`) would inevitably cause import cycles if resources were split into sub-packages without a shared types layer.

---

## Target Architecture

### 1. Shared REST Infrastructure: `internal/restapi`

A cross-system package that defines the common HTTP contract and a reusable base implementation.

```
internal/restapi/
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

### 2. Shared Models: `internal/netbox/models`

All data structs live in a single, cycle-free package. Sub-packages (`dcim/`, `ipam/`) import `models` but never each other.

```
internal/netbox/
  models/
    base.go        # ApiBaseFields, ApiListResponse[T]
    dcim.go        # Device, Site, DeviceType, DeviceRole, Cable...
    ipam.go        # Prefix, IPAddress, VRF, VLAN...
    circuits.go    # Circuit, Provider...
    virtualization.go
```

### 3. System-Specific API Packages

Each domain gets its own package. Resource logic lives in sub-packages that depend only on `restapi.Doer` and `models`.

```
internal/netbox/
  client.go          # NetboxClient: embeds *restapi.BaseClient, wires resource APIs
  aliases.go         # Type aliases so callers still see netbox.Device, netbox.Site, etc.

  dcim/
    device.go        # DeviceAPI
    cable.go         # CableAPI
    site.go          # SiteAPI
    device_type.go   # DeviceTypeAPI
    device_role.go   # DeviceRoleAPI
    rack.go          # ... etc
    ...

  ipam/
    prefix.go        # PrefixAPI
    ipaddress.go     # IPAddressAPI
    vrf.go           # VRFAPI
    ...

  circuits/
    ...

  tenancy/
    ...

  virtualization/
    ...

internal/servicenow/
  client.go          # Also embeds *restapi.BaseClient
  auth.go            # ServiceNow-specific auth
  incident/
    ...
```

### 4. Flat API Surface

Drop the `DCIM` / `IPAM` grouping. Resource names are distinctive enough.

```go
client, _ := netbox.NewClient()
client.Device.List(ctx)
client.Site.Get(ctx, 42)
client.Prefix.Create(ctx, &prefix)
```

### 5. Why This Avoids Import Cycles

- `restapi` → depends on nothing
- `models` → depends on nothing
- `netbox/dcim` → depends on `restapi` + `models`
- `netbox/ipam` → depends on `restapi` + `models`
- `netbox` (root) → depends on `restapi` + `models` + `netbox/dcim` + `netbox/ipam`
- No package imports its parent. Clean DAG.

Cross-domain references (e.g. `Device.PrimaryIP4` referencing `ipam.IPAddress`) are safe because both domains import the same `models` package instead of importing each other.

---

## Tradeoffs and Decisions

### Decision: Package name is `internal/restapi`

**Rationale:** It signals "common REST client infrastructure" without being as low-level as raw `transport`. Alternatives considered: `transport` (sounds like TCP/TLS), `client` (too generic, clashes with `netbox.Client`), `httpkit` (accurate but verbose). `restapi` is concise and intent-revealing.

### Decision: Include `BaseClient` in `restapi`, not just the interface

**Rationale:** Netbox and ServiceNow clients become tiny (~20 lines). All HTTP boilerplate lives in one place. The interface alone would force copy-pasting the same `http.NewRequestWithContext`, header merging, and error formatting in every new system. If an exotic client (gRPC, heavy middleware) appears later, we can introduce a second interface — the cost of that future flexibility today is too high.

### Decision: Auth plugs into `BaseClient` at construction time

**Rationale:** No separate `Auth()` method. The Netbox `NewClient()` loads the API key from env and injects it into `BaseClient.Header`. ServiceNow will do the same with its own mechanism (basic auth token, OAuth bearer, etc.). Auth becomes a configuration detail, not a runtime state mutation.

### Decision: Generic response helpers live in `restapi`, list response types stay in `netbox/models`

**Rationale:** `IsSuccess`, `IsError`, and body-reading error formatting are 100% generic. `ApiListResponse[T]` is Netbox-specific (`count`, `next`, `previous`, `results`). ServiceNow may paginate differently, so list wrappers stay local to the system package unless proven reusable.

### Decision: All model structs live in `internal/netbox/models`

**Rationale:** This is the only way to support cross-domain references (e.g. `Device` → `IPAddress`) without circular imports. Sub-packages (`dcim/`, `ipam/`) import `models` but never each other. A future ServiceNow client would have its own `internal/servicenow/models` rather than sharing Netbox types.

### Decision: Root package re-exports types via aliases

**Rationale:** Callers today use a single import (`internal/netbox`) and refer to `netbox.Device`. Re-exporting (`type Device = models.Device`) preserves this ergonomics without requiring callers to import `models` directly. Under the hood there is zero runtime cost.

### Decision: API struct names are singular

**Rationale:** Consistency. `DeviceAPI`, `SiteAPI`, `DeviceTypeAPI`, `DeviceRoleAPI`, `CableAPI`, `PrefixAPI`, etc. The old code mixed singular and plural (`SitesAPI`, `DeviceTypesAPI`, `DeviceRolesAPI`); standardizing removes guesswork.

---

## Implementation Steps

### Phase 1: Bootstrap `internal/restapi`

1. Create `internal/restapi/doer.go` with the `Doer` interface.
2. Create `internal/restapi/client.go` with `BaseClient`:
   - Fields: `BaseURL string`, `HTTP *http.Client`, `Header http.Header`
   - Methods implementing `Doer`
3. Create `internal/restapi/errors.go` with:
   - `IsSuccess(res *http.Response) bool`
   - `IsError(res *http.Response) bool`
   - `APIError(res *http.Response) error` — body-reading, HTML guard, truncation logic extracted from current `NetboxClient.apiError()`

### Phase 2: Create `internal/netbox/models`

1. Create `internal/netbox/models/base.go` with `ApiBaseFields` and `ApiListResponse[T]`.
2. Create `internal/netbox/models/dcim.go` with all DCIM structs:
   - `Device`, `DeviceWrite`, `(*Device).MapToJsonWrite()`
   - `Site`, `SiteWrite`
   - `DeviceType`, `DeviceTypeWrite`
   - `DeviceRole`, `DeviceRoleWrite`
   - `Cable`, `CableWrite`
3. Create `internal/netbox/models/ipam.go` with IPAM skeleton structs (e.g. `Prefix`, `PrefixWrite`).

### Phase 3: Restructure `internal/netbox`

1. Create `internal/netbox/dcim/`.
   - Move full API logic → `dcim/device.go`, `dcim/site.go`, `dcim/device_type.go`, `dcim/device_role.go`, `dcim/cable.go`.
   - Change `*API` struct field from `*NetboxClient` to `restapi.Doer`.
   - Constructors: `NewDeviceAPI(d restapi.Doer) *DeviceAPI`, etc.
   - Use `restapi.IsError` and `restapi.APIError` instead of current helpers.
   - Use `models.ApiListResponse[T]` and `models.Device`, etc.
2. Create `internal/netbox/ipam/`.
   - Move/create future IPAM resource APIs.
3. Rewrite `internal/netbox/client.go`.
   - `NetboxClient` embeds `*restapi.BaseClient`.
   - Remove `get`, `post`, `patch`, `delete` methods (now handled by `BaseClient`).
   - Flatten API surface: `Device`, `Cable`, `Site`, `DeviceType`, `DeviceRole`, `Prefix`, etc. directly on `NetboxClient`.
   - `NewClient()` loads env vars, creates `BaseClient`, sets auth headers, then constructs each resource API with `dcim.NewDeviceAPI(c)`, etc.
4. Create `internal/netbox/aliases.go`.
   - Re-export `models.Device`, `models.Site`, `models.Prefix`, etc. as `netbox.Device`, `netbox.Site`, `netbox.Prefix`.
   - Re-export `models.ApiListResponse[T]` and `models.ApiBaseFields`.
5. Delete old files: `api.go`, `auth.go`, `dcim.go`, `ipam.go`, `devices.go`, `cables.go`, `sites.go`, `device_types.go`, `device_roles.go`.

### Phase 4: Update Callers

1. Update `cmd/rundeck/netbox/clone/main.go` and any other `cmd/` files.
   - Change `client.DCIM.Device.List(ctx)` → `client.Device.List(ctx)`.
   - No import path changes needed thanks to root-package type aliases.

### Phase 5: Scale to Full API

1. For each new Netbox resource:
   - Add its struct(s) to the appropriate `models/` file (or create a new one if the domain doesn't exist yet).
   - Create a file in the appropriate sub-package (`dcim/`, `ipam/`, etc.).
   - Define `*API` struct.
   - Implement CRUD methods using `restapi.Doer`.
   - Wire constructor in `NetboxClient`.

### Phase 6 (Optional): Generic CRUD Helper

If the repetition of `List` / `Get` / `Create` / `Update` / `Delete` across 50+ resources becomes tedious:

```go
// In restapi or a helper inside netbox
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
    *restapi.ResourceAPI[models.Device]
}

func NewDeviceAPI(d restapi.Doer) *DeviceAPI {
    return &DeviceAPI{restapi.NewResourceAPI[models.Device](d, "/api/dcim/devices/")}
}
```

Evaluate after Phase 5 if the boilerplate justifies the abstraction.

---

## Future Systems

When adding ServiceNow (or any other system):

1. Create `internal/<system>/`.
2. Embed `*restapi.BaseClient` in `<system>.Client`.
3. Write `<system>/auth.go` for its specific auth mechanism.
4. Create `internal/<system>/models/` for its data structs.
5. Create sub-packages for its resource domains.
6. Resource APIs depend only on `restapi.Doer` — zero HTTP code duplication.

---

## Checklist

- [ ] Phase 1: `internal/restapi` created and compiles.
- [ ] Phase 2: `internal/netbox/models` created with base + DCIM + IPAM skeletons.
- [ ] Phase 3: `internal/netbox` restructured into sub-packages with type aliases.
- [ ] Phase 4: All `cmd/` callers updated and working.
- [ ] Phase 5: Existing resources (Device, Cable, Site, DeviceType, DeviceRole) ported.
- [ ] Phase 6 (optional): Generic `ResourceAPI[T]` evaluated and optionally implemented.
