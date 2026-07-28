---
title: waitingroom_query_event
page_id: schema-waitingroom-query-event-04ce8a11
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_query_event

```yaml
{"type": "object", "properties": {"custom_page_html": {"$ref": "#/components/schemas/waitingroom_event_custom_page_html"}, "description": {"$ref": "#/components/schemas/waitingroom_event_description"}, "disable_session_renewal": {"$ref": "#/components/schemas/waitingroom_event_disable_session_renewal"}, "event_end_time": {"$ref": "#/components/schemas/waitingroom_event_end_time"}, "event_start_time": {"$ref": "#/components/schemas/waitingroom_event_start_time"}, "name": {"$ref": "#/components/schemas/waitingroom_event_name"}, "new_users_per_minute": {"$ref": "#/components/schemas/waitingroom_event_new_users_per_minute"}, "prequeue_start_time": {"$ref": "#/components/schemas/waitingroom_event_prequeue_start_time"}, "queueing_method": {"$ref": "#/components/schemas/waitingroom_event_queueing_method"}, "session_duration": {"$ref": "#/components/schemas/waitingroom_event_session_duration"}, "shuffle_at_event_start": {"$ref": "#/components/schemas/waitingroom_event_shuffle_at_event_start"}, "suspended": {"$ref": "#/components/schemas/waitingroom_event_suspended"}, "total_active_users": {"$ref": "#/components/schemas/waitingroom_event_total_active_users"}, "turnstile_action": {"$ref": "#/components/schemas/waitingroom_event_turnstile_action"}, "turnstile_mode": {"$ref": "#/components/schemas/waitingroom_event_turnstile_mode"}}, "required": ["name", "event_start_time", "event_end_time"]}
```
