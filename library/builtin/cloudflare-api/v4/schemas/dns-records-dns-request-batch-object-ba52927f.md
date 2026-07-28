---
title: dns-records_dns-request-batch-object
page_id: schema-dns-records-dns-request-batch-object-ba52927f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns-request-batch-object

```yaml
{"type": "object", "properties": {"deletes": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-batch-delete"}}, "patches": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-batch-patch"}}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-batch-post"}}, "puts": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-batch-put"}}}}
```
