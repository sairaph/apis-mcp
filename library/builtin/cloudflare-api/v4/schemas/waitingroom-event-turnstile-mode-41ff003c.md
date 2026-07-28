---
title: waitingroom_event_turnstile_mode
page_id: schema-waitingroom-event-turnstile-mode-41ff003c
path: schemas
description: If set, the event will override the waiting room's `turnstile_mode` property while it is active. If null, the event will inherit it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_event_turnstile_mode

If set, the event will override the waiting room's `turnstile_mode` property while it is active. If null, the event will inherit it.

```yaml
{"description": "If set, the event will override the waiting room's `turnstile_mode` property while it is active. If null, the event will inherit it.", "type": "string", "enum": ["off", "invisible", "visible_non_interactive", "visible_managed"], "nullable": true, "x-auditable": true}
```
