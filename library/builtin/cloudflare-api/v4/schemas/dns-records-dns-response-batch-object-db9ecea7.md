---
title: dns-records_dns-response-batch-object
page_id: schema-dns-records-dns-response-batch-object-db9ecea7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns-response-batch-object

```yaml
{"type": "object", "properties": {"deletes": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-response"}}, "patches": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-response"}}, "posts": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-response"}}, "puts": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-response"}}}}
```
