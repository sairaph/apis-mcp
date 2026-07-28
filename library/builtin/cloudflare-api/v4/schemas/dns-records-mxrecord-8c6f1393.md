---
title: dns-records_MXRecord
page_id: schema-dns-records-mxrecord-8c6f1393
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_MXRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A valid mail server hostname.", "example": "mx.example.com", "format": "hostname", "type": "string", "x-auditable": true}, "priority": {"$ref": "#/components/schemas/dns-records_priority"}, "type": {"description": "Record type.", "type": "string", "example": "MX", "enum": ["MX"], "x-auditable": true}}, "type": "object"}], "title": "MX Record"}
```
