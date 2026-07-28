---
title: mcn_onramp_status
page_id: schema-mcn-onramp-status-9abbff74
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_onramp_status

```yaml
{"type": "object", "properties": {"apply_progress": {"$ref": "#/components/schemas/mcn_apply_progress"}, "lifecycle_errors": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/mcn_error"}}, "lifecycle_state": {"$ref": "#/components/schemas/mcn_onramp_lifecycle_state"}, "plan_progress": {"$ref": "#/components/schemas/mcn_plan_progress"}, "routes": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_conduit_route_id"}}, "tunnels": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_conduit_tunnel_id"}}}, "required": ["lifecycle_state", "tunnels", "routes", "plan_progress", "apply_progress"]}
```
