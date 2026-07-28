---
title: dns-records_TLSARecord
page_id: schema-dns-records-tlsarecord-8c660f4e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_TLSARecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted TLSA content. See 'data' to set TLSA properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a TLSA record.", "type": "object", "properties": {"certificate": {"description": "Certificate.", "type": "string", "x-auditable": true}, "matching_type": {"description": "Matching Type.", "type": "number", "example": 1, "maximum": 255, "minimum": 0, "x-auditable": true}, "selector": {"description": "Selector.", "type": "number", "example": 0, "maximum": 255, "minimum": 0, "x-auditable": true}, "usage": {"description": "Usage.", "type": "number", "example": 0, "maximum": 255, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "TLSA", "enum": ["TLSA"], "x-auditable": true}}, "type": "object"}], "title": "TLSA Record"}
```
