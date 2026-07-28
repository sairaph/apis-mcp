---
title: teams-devices_tanium_input_request
page_id: schema-teams-devices-tanium-input-request-8404f700
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_tanium_input_request

```yaml
{"type": "object", "properties": {"connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f", "x-auditable": true}, "eid_last_seen": {"description": "For more details on eid last seen, refer to the Tanium documentation.", "type": "string", "example": "2023-07-20T23:16:32Z", "x-auditable": true}, "operator": {"description": "Operator to evaluate risk_level or eid_last_seen.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}, "risk_level": {"description": "For more details on risk level, refer to the Tanium documentation.", "type": "string", "example": "low", "enum": ["low", "medium", "high", "critical"], "x-auditable": true}, "scoreOperator": {"description": "Score Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}, "total_score": {"description": "For more details on total score, refer to the Tanium documentation.", "type": "number", "example": 1, "x-auditable": true}}, "required": ["connection_id"], "title": "Tanium S2S Input"}
```
