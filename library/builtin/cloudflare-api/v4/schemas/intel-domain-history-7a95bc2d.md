---
title: intel_domain-history
page_id: schema-intel-domain-history-7a95bc2d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_domain-history

```yaml
{"type": "object", "properties": {"categorizations": {"type": "array", "items": {"properties": {"categories": {"type": "array", "items": {"properties": {"id": {"type": "integer", "example": 155}, "name": {"type": "string", "example": "Technology"}}, "type": "object"}, "example": [{"id": 155, "name": "Technology"}]}, "end": {"type": "string", "format": "date", "example": "2021-04-30", "x-auditable": true}, "start": {"type": "string", "format": "date", "example": "2021-04-01", "x-auditable": true}}, "type": "object"}}, "domain": {"$ref": "#/components/schemas/intel_domain_name"}}}
```
