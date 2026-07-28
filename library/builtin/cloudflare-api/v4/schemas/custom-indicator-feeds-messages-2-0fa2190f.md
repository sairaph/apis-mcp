---
title: custom-indicator-feeds_messages-2
page_id: schema-custom-indicator-feeds-messages-2-0fa2190f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_messages-2

```yaml
{"type": "array", "items": {"properties": {"code": {"type": "integer", "minimum": 1000}, "documentation_url": {"type": "string"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object", "uniqueItems": true}, "example": []}
```
