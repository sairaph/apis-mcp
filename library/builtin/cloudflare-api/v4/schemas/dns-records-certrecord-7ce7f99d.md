---
title: dns-records_CERTRecord
page_id: schema-dns-records-certrecord-7ce7f99d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_CERTRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted CERT content. See 'data' to set CERT properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a CERT record.", "type": "object", "properties": {"algorithm": {"description": "Algorithm.", "type": "number", "example": 8, "maximum": 255, "minimum": 0, "x-auditable": true}, "certificate": {"description": "Certificate.", "type": "string", "x-auditable": true}, "key_tag": {"description": "Key Tag.", "type": "number", "example": 1, "maximum": 65535, "minimum": 0, "x-auditable": true}, "type": {"description": "Type.", "type": "number", "example": 9, "maximum": 65535, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "CERT", "enum": ["CERT"], "x-auditable": true}}, "type": "object"}], "title": "CERT Record"}
```
