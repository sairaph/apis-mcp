---
title: dns-records_PTRRecord
page_id: schema-dns-records-ptrrecord-223ee127
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_PTRRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Domain name pointing to the address.", "example": "example.com", "type": "string", "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "PTR", "enum": ["PTR"], "x-auditable": true}}, "type": "object"}], "title": "PTR Record"}
```
