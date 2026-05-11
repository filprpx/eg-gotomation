# NetBox Work Handover

## Current state

The codebase now supports two interaction styles through `internal/netbox`:

1. Typed workflows through `internal/netbox/models` and typed resource APIs.
2. Untyped workflows through request-based transport and raw JSON access.

The transport layer in `internal/restapi` was simplified around a request object.

## Transport and API architecture

### `internal/restapi`

- `Doer` is request-based:
  - `Do(ctx, Request) (*http.Response, error)`
- `Request` carries:
  - `Method`
  - `Path`
  - `Query`
  - `Body`
- `Client` is the concrete transport implementation.
- `DefaultConfig()` replaced the old `NewConfig()` naming.

### `internal/netbox/api`

- `BaseAPI[T]` uses `restapi.Doer` and request objects.
- `RawAPI` exists for untyped access and returns `map[string]any` / `[]map[string]any`.

### `internal/netbox/client.go`

- Embeds `restapi.Client`
- Exposes typed resource APIs
- Exposes `Raw *api.RawAPI`

## Modeling conventions settled so far

For a resource `X`:

1. `NestedX`
   - shape used when the resource is nested inside another payload.

2. `X`
   - full read model.
   - usually embeds `NestedX`.

3. `xCommon`
   - only fields that:
     - are writable
     - have the same type in read and write
     - are not already present in `APIBaseFields` or `NestedX`

4. `XWrite`
   - write payload model.
   - embeds `APIWriteBaseFields` and `xCommon`.
   - explicitly includes:
     - ID-mapped nested references
     - read `Choice` -> write primitive conversions
     - `Comments` when supported, because `APIWriteBaseFields` does not contain it.

## Important helper decisions

### `Choice`

`Choice.Value` was widened to `any` because NetBox returns both string and numeric choice values.

Helpers in `internal/netbox/models/base.go`:

- `safeChoiceValue`
- `safeChoiceIntValue`
- `SafeGetId`

### Tags

Tags are still modeled with `[]NestedTag`-style structures.

This is likely good enough for current work, but write behavior should be revalidated later against real NetBox requests. A future cleanup may normalize tag write payloads globally.

### Comments

`Comments` exists in `APIBaseFields` but not in `APIWriteBaseFields`, so every write model that supports comments must declare `Comments` explicitly.

## Completed model batches

### Batch A

- `Region`
- `SiteGroup`
- `RackRole`
- `Manufacturer`
- `Platform`
- `DeviceRole`

### Batch B

- `Site`
- `Location`
- `Rack`

Batch B was first applied conservatively, then corrected from real API samples.

### Batch C

- `Tenant`
- `IPAddress`
- `Cluster`
- `DeviceType` refinement

Supporting nested types added during this work:

- `ASN`
- `RIR`
- `VRF`
- `TenantGroup`
- `ClusterType`
- `ClusterGroup`
- `ConfigTemplate`

## Files notably touched

### Transport / API

- `internal/restapi/doer.go`
- `internal/restapi/client.go`
- `internal/restapi/config.go`
- `internal/netbox/api/base.go`
- `internal/netbox/api/raw.go`
- `internal/netbox/client.go`
- `internal/netbox/config.go`

### Models

- `internal/netbox/models/base.go`
- `internal/netbox/models/device.go`
- `internal/netbox/models/device_role.go`
- `internal/netbox/models/device_type.go`
- `internal/netbox/models/location.go`
- `internal/netbox/models/manufacturer.go`
- `internal/netbox/models/site.go`
- `internal/netbox/models/tenant.go`
- plus all newly added support model files in `internal/netbox/models/`

## Remaining functional work

### 1. `Cable`

This is the main remaining original model area that still needs accurate schema-driven work.

### 2. Consistency audit

Recommended before expanding much further:

- Re-review `Device` now that supporting nested models are richer.
- Check that each `NestedX` contains everything shown in real nested payloads.
- Verify comments and write mapping consistency across all current models.
- Revalidate tag write behavior.

### 3. Optional full-resource completion

Only if typed CRUD is wanted for them as full resources:

- `ASN`
- `VRF`
- `TenantGroup`
- `ClusterType`
- `ClusterGroup`
- `RIR`

## Streamlining the schema extraction process

Manual Swagger/OpenAPI inspection is currently the slowest part of the workflow.

### Goal

Build a helper that, for a resource, prints:

- GET item fields
- POST body fields
- read/write type changes
- nested refs implying `NestedX`
- read-only fields and counters
- candidate `xCommon` fields

### Suggested helper behavior

Input:

`netbox-schema inspect dcim/sites`

Output:

- fields present in GET item payload
- fields present in POST payload
- type transitions like:
  - `status: Choice -> string`
  - `tenant: object -> int`
  - `width: Choice -> int`
- likely `NestedX` references
- likely read-only fields

### Suggested implementation strategy

1. Use the NetBox OpenAPI schema as primary source.
2. Resolve GET/POST resource schemas.
3. Diff read and write shapes.
4. Print a human-readable worksheet, not generated Go code.

### Why not full code generation

The current hand-curated model pattern is better than autogenerated models for this repo because:

- read and write shapes differ intentionally
- nested models are curated, not purely mechanical
- the goal is maintainable modeling, not broad autogenerated coverage

## Recommended next session

1. Review `Cable` from real GET/POST samples.
2. Run a consistency audit on `Device` and the completed batches.
3. Start the schema-inspection helper instead of continuing fully manual Swagger extraction.

## Git notes

The work was intended to be split conceptually into:

1. transport/raw-access compatibility changes
2. typed model expansion and cleanup

If continuing in a fresh session, start by reading:

- `internal/restapi/doer.go`
- `internal/netbox/api/base.go`
- `internal/netbox/api/raw.go`
- `internal/netbox/client.go`
- `internal/netbox/models/base.go`
- `internal/netbox/models/device.go`
