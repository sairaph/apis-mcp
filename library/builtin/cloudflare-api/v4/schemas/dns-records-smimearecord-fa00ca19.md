---
title: dns-records_SMIMEARecord
page_id: schema-dns-records-smimearecord-fa00ca19
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_SMIMEARecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted SMIMEA content. See 'data' to set SMIMEA properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a SMIMEA record.", "type": "object", "properties": {"certificate": {"description": "Certificate.", "type": "string", "x-auditable": true}, "matching_type": {"description": "Matching Type.", "type": "number", "example": 0, "maximum": 255, "minimum": 0, "x-auditable": true}, "selector": {"description": "Selector.", "type": "number", "example": 0, "maximum": 255, "minimum": 0, "x-auditable": true}, "usage": {"description": "Usage.", "type": "number", "example": 3, "maximum": 255, "minimum": 0, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "SMIMEA", "enum": ["SMIMEA"], "x-auditable": true}}, "type": "object"}], "title": "SMIMEA Record"}
```
