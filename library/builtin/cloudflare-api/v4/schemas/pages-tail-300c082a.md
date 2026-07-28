---
title: pages_tail
page_id: schema-pages-tail-300c082a
path: schemas
description: A tail session for streaming logs from a Pages deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_tail

A tail session for streaming logs from a Pages deployment.

```yaml
{"description": "A tail session for streaming logs from a Pages deployment.", "type": "object", "properties": {"id": {"description": "Identifier of the tail session.", "type": "string", "example": "49a4dcf81a3940fab8453b2be3fb86ef", "readOnly": true}, "url": {"description": "Optional WebSocket URL to connect to for receiving tail events, when returned by the tail service.", "type": "string", "example": "wss://tail.developers.workers.dev/49a4dcf81a3940fab8453b2be3fb86ef", "readOnly": true}}, "required": ["id"]}
```
