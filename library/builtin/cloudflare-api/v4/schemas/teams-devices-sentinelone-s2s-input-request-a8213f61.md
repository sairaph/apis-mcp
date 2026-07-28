---
title: teams-devices_sentinelone_s2s_input_request
page_id: schema-teams-devices-sentinelone-s2s-input-request-a8213f61
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_sentinelone_s2s_input_request

```yaml
{"type": "object", "properties": {"active_threats": {"description": "The Number of active threats.", "type": "number", "example": 1}, "connection_id": {"description": "Posture Integration ID.", "type": "string", "example": "bc7cbfbb-600a-42e4-8a23-45b5e85f804f"}, "infected": {"description": "Whether device is infected.", "type": "boolean", "example": true}, "is_active": {"description": "Whether device is active.", "type": "boolean", "example": true}, "network_status": {"description": "Network status of device.", "type": "string", "example": "connected", "enum": ["connected", "disconnected", "disconnecting", "connecting"]}, "operational_state": {"description": "Agent operational state.", "type": "string", "enum": ["na", "partially_disabled", "auto_fully_disabled", "fully_disabled", "auto_partially_disabled", "disabled_error", "db_corruption"]}, "operator": {"description": "Operator.", "type": "string", "example": ">", "enum": ["<", "<=", ">", ">=", "=="]}}, "required": ["connection_id"], "title": "SentinelOne S2S Input"}
```
