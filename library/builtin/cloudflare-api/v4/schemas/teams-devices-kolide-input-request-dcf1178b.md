---
title: teams-devices_kolide_input_request
page_id: schema-teams-devices-kolide-input-request-dcf1178b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_kolide_input_request

```yaml
{"type": "object", "properties": {"auth_state": {"description": "The set of Kolide device authentication states that pass the posture check. Device must match one of the specified states.", "type": "array", "items": {"enum": ["Good", "Notified", "Will Block", "Blocked"], "type": "string"}, "example": ["Good", "Notified"], "minItems": 1, "uniqueItems": true, "x-auditable": true}, "connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f", "x-auditable": true}, "countOperator": {"description": "Count Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}, "issue_count": {"description": "The Number of Issues.", "type": "string", "example": "1", "x-auditable": true}}, "required": ["connection_id"], "title": "Kolide S2S Input"}
```
