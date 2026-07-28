---
title: moq_relay
page_id: schema-moq-relay-f88aa9bc
path: schemas
description: Full relay details (no tokens).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_relay

Full relay details (no tokens).

```yaml
{"description": "Full relay details (no tokens).", "type": "object", "properties": {"config": {"$ref": "#/components/schemas/moq_relay_config"}, "created": {"type": "string", "format": "date-time"}, "modified": {"type": "string", "format": "date-time"}, "name": {"type": "string", "example": "Production Live Stream"}, "status": {"description": "\"connected\" when active, omitted otherwise.", "type": "string", "enum": ["connected"]}, "uid": {"type": "string", "example": "a1b2c3d4e5f67890a1b2c3d4e5f67890"}}, "required": ["uid", "created", "modified", "name", "config"]}
```
