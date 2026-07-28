---
title: dns-records_NSRecord
page_id: schema-dns-records-nsrecord-605f6bd3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_NSRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A valid name server host name.", "example": "ns1.example.com", "type": "string", "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "NS", "enum": ["NS"], "x-auditable": true}}, "type": "object"}], "title": "NS Record"}
```
