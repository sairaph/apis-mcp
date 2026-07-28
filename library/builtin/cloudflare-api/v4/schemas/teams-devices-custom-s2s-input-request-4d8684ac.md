---
title: teams-devices_custom_s2s_input_request
page_id: schema-teams-devices-custom-s2s-input-request-4d8684ac
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_custom_s2s_input_request

```yaml
{"type": "object", "properties": {"connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f"}, "operator": {"description": "Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="]}, "score": {"description": "A value between 0-100 assigned to devices set by the 3rd party posture provider.", "type": "number", "example": 100}}, "required": ["connection_id", "score", "operator"], "title": "Custom Device Posture Integration Input"}
```
