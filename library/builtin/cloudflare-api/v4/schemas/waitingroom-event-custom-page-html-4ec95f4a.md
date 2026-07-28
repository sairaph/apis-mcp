---
title: waitingroom_event_custom_page_html
page_id: schema-waitingroom-event-custom-page-html-4ec95f4a
path: schemas
description: If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_event_custom_page_html

If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.

```yaml
{"description": "If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.", "type": "string", "example": "{{#waitTimeKnown}} {{waitTime}} mins {{/waitTimeKnown}} {{^waitTimeKnown}} Event is prequeueing / Queue all enabled {{/waitTimeKnown}}", "nullable": true, "x-auditable": true}
```
