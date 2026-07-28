---
title: waitingroom_event_details_result
page_id: schema-waitingroom-event-details-result-254f6bab
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_event_details_result

```yaml
{"type": "object", "properties": {"created_on": {"$ref": "#/components/schemas/waitingroom_timestamp"}, "custom_page_html": {"$ref": "#/components/schemas/waitingroom_event_details_custom_page_html"}, "description": {"$ref": "#/components/schemas/waitingroom_event_description"}, "disable_session_renewal": {"$ref": "#/components/schemas/waitingroom_event_details_disable_session_renewal"}, "event_end_time": {"$ref": "#/components/schemas/waitingroom_event_end_time"}, "event_start_time": {"$ref": "#/components/schemas/waitingroom_event_start_time"}, "id": {"$ref": "#/components/schemas/waitingroom_event_id"}, "modified_on": {"$ref": "#/components/schemas/waitingroom_timestamp"}, "name": {"$ref": "#/components/schemas/waitingroom_event_name"}, "new_users_per_minute": {"$ref": "#/components/schemas/waitingroom_event_details_new_users_per_minute"}, "prequeue_start_time": {"$ref": "#/components/schemas/waitingroom_event_prequeue_start_time"}, "queueing_method": {"$ref": "#/components/schemas/waitingroom_event_details_queueing_method"}, "session_duration": {"$ref": "#/components/schemas/waitingroom_event_details_session_duration"}, "shuffle_at_event_start": {"$ref": "#/components/schemas/waitingroom_event_shuffle_at_event_start"}, "suspended": {"$ref": "#/components/schemas/waitingroom_event_suspended"}, "total_active_users": {"$ref": "#/components/schemas/waitingroom_event_details_total_active_users"}}}
```
