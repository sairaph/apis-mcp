---
title: secondary-dns_messages
page_id: schema-secondary-dns-messages-130b5c40
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_messages

```yaml
{"type": "array", "items": {"properties": {"code": {"type": "integer", "minimum": 1000}, "documentation_url": {"type": "string"}, "message": {"type": "string"}, "source": {"type": "object", "properties": {"pointer": {"type": "string"}}}}, "required": ["code", "message"], "type": "object", "uniqueItems": true}, "example": []}
```
