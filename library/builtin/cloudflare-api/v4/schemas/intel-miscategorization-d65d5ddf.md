---
title: intel_miscategorization
page_id: schema-intel-miscategorization-d65d5ddf
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_miscategorization

```yaml
{"type": "object", "properties": {"content_adds": {"description": "Content category IDs to add.", "type": "array", "items": {"type": "integer", "x-auditable": true}, "example": [82]}, "content_removes": {"description": "Content category IDs to remove.", "type": "array", "items": {"type": "integer", "x-auditable": true}, "example": [155]}, "indicator_type": {"type": "string", "example": "domain", "enum": ["domain", "ipv4", "ipv6", "url"], "x-auditable": true}, "ip": {"description": "Provide only if indicator_type is `ipv4` or `ipv6`.", "type": "string", "nullable": true, "x-auditable": true}, "security_adds": {"description": "Security category IDs to add.", "type": "array", "items": {"type": "integer", "x-auditable": true}, "example": [117, 131]}, "security_removes": {"description": "Security category IDs to remove.", "type": "array", "items": {"type": "integer", "x-auditable": true}, "example": [83]}, "url": {"description": "Provide only if indicator_type is `domain` or `url`. Example if indicator_type is `domain`: `example.com`. Example if indicator_type is `url`: `https://example.com/news/`.", "type": "string", "x-auditable": true}}}
```
