---
title: dns-records_NAPTRRecord
page_id: schema-dns-records-naptrrecord-c0f4849b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_NAPTRRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted NAPTR content. See 'data' to set NAPTR properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a NAPTR record.", "type": "object", "properties": {"flags": {"description": "Flags.", "type": "string", "x-auditable": true}, "order": {"description": "Order.", "type": "number", "example": 100, "maximum": 65535, "minimum": 0, "x-auditable": true}, "preference": {"description": "Preference.", "type": "number", "example": 10, "maximum": 65535, "minimum": 0, "x-auditable": true}, "regex": {"description": "Regex.", "type": "string", "x-auditable": true}, "replacement": {"description": "Replacement.", "type": "string", "x-auditable": true}, "service": {"description": "Service.", "type": "string", "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "NAPTR", "enum": ["NAPTR"], "x-auditable": true}}, "type": "object"}], "title": "NAPTR Record"}
```
