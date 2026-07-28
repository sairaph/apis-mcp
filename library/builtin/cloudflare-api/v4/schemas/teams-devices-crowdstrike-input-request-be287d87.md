---
title: teams-devices_crowdstrike_input_request
page_id: schema-teams-devices-crowdstrike-input-request-be287d87
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_crowdstrike_input_request

```yaml
{"type": "object", "properties": {"connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f", "x-auditable": true}, "last_seen": {"description": "For more details on last seen, please refer to the Crowdstrike documentation.", "type": "string", "example": "15d3h20m4s", "x-auditable": true}, "operator": {"description": "Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}, "os": {"description": "Os Version.", "type": "string", "example": "13.3.0", "x-auditable": true}, "overall": {"description": "Overall.", "type": "string", "example": 90, "x-auditable": true}, "sensor_config": {"description": "SensorConfig.", "type": "string", "example": 90, "x-auditable": true}, "state": {"description": "For more details on state, please refer to the Crowdstrike documentation.", "type": "string", "example": "online", "enum": ["online", "offline", "unknown"], "x-auditable": true}, "version": {"description": "Version.", "type": "string", "example": "13.3.0", "x-auditable": true}, "versionOperator": {"description": "Version Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}}, "required": ["connection_id"], "title": "Crowdstrike S2S Input"}
```
