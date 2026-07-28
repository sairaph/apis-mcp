---
title: dns-records_dns_response_collection
page_id: schema-dns-records-dns-response-collection-6f2f1e50
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_dns_response_collection

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/dns-records_dns-record-response"}}}, "type": "object"}]}
```
