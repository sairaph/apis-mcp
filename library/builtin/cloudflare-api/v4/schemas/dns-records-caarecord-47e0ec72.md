---
title: dns-records_CAARecord
page_id: schema-dns-records-caarecord-47e0ec72
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_CAARecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted CAA content. See 'data' to set CAA properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a CAA record.", "type": "object", "properties": {"flags": {"description": "Flags for the CAA record.", "type": "number", "example": 1, "maximum": 255, "minimum": 0, "x-auditable": true}, "tag": {"description": "Name of the property controlled by this record (e.g.: issue, issuewild, iodef).", "type": "string", "example": "issue", "x-auditable": true}, "value": {"description": "Value of the record. This field's semantics depend on the chosen tag.", "type": "string", "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "CAA", "enum": ["CAA"], "x-auditable": true}}, "type": "object"}], "title": "CAA Record"}
```
