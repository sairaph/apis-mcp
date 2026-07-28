---
title: waitingroom_event_turnstile_action
page_id: schema-waitingroom-event-turnstile-action-09fee2d2
path: schemas
description: If set, the event will override the waiting room's `turnstile_action` property while it is active. If null, the event will inherit it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_event_turnstile_action

If set, the event will override the waiting room's `turnstile_action` property while it is active. If null, the event will inherit it.

```yaml
{"description": "If set, the event will override the waiting room's `turnstile_action` property while it is active. If null, the event will inherit it.", "type": "string", "enum": ["log", "infinite_queue"], "nullable": true, "x-auditable": true}
```
