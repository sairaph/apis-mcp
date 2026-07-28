---
title: intel_schemas-asn
page_id: schema-intel-schemas-asn-13e46e45
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_schemas-asn

```yaml
{"type": "object", "properties": {"asn": {"$ref": "#/components/schemas/intel_asn"}, "country": {"$ref": "#/components/schemas/intel_asn_country"}, "description": {"$ref": "#/components/schemas/intel_asn_description"}, "domain_count": {"type": "integer", "x-auditable": true}, "top_domains": {"type": "array", "items": {"type": "string", "x-auditable": true}, "example": ["example.com"]}, "type": {"$ref": "#/components/schemas/intel_asn_type"}}}
```
