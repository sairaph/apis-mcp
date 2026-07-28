---
title: waitingroom_status_response
page_id: schema-waitingroom-status-response-812b6d12
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_status_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/waitingroom_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"estimated_queued_users": {"$ref": "#/components/schemas/waitingroom_estimated_queued_users"}, "estimated_total_active_users": {"$ref": "#/components/schemas/waitingroom_estimated_total_active_users"}, "event_id": {"$ref": "#/components/schemas/waitingroom_status_event_id"}, "max_estimated_time_minutes": {"$ref": "#/components/schemas/waitingroom_max_estimated_time_minutes"}, "status": {"$ref": "#/components/schemas/waitingroom_status"}}}}, "type": "object"}]}
```
