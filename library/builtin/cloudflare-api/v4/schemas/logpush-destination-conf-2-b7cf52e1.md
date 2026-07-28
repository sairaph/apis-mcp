---
title: logpush_destination_conf-2
page_id: schema-logpush-destination-conf-2-b7cf52e1
path: schemas
description: Unique WebSocket address that will receive messages from Cloudflare’s edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_destination_conf-2

Unique WebSocket address that will receive messages from Cloudflare’s edge.

```yaml
{"description": "Unique WebSocket address that will receive messages from Cloudflare’s edge.", "type": "string", "format": "uri", "example": "wss://logs.cloudflare.com/instant-logs/ws/sessions/99d471b1ca3c23cc8e30b6acec5db987", "maxLength": 4096, "x-auditable": true}
```
