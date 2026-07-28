---
title: waitingroom_event_session_duration
page_id: schema-waitingroom-event-session-duration-d21d4ade
path: schemas
description: If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_event_session_duration

If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.

```yaml
{"description": "If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.", "type": "integer", "maximum": 30, "minimum": 1, "nullable": true, "x-auditable": true}
```
