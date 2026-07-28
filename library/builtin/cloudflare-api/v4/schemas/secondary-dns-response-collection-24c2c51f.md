---
title: secondary-dns_response_collection
page_id: schema-secondary-dns-response-collection-24c2c51f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_response_collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/secondary-dns_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/secondary-dns_tsig"}}}, "type": "object"}]}
```
