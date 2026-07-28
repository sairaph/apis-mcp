---
title: dns-records_SVCBRecord
page_id: schema-dns-records-svcbrecord-6bf59281
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_SVCBRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted SVCB content. See 'data' to set SVCB properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a SVCB record.", "type": "object", "properties": {"priority": {"description": "Priority.", "type": "number", "example": 1, "maximum": 65535, "minimum": 0, "x-auditable": true}, "target": {"description": "Target.", "type": "string", "example": ".", "x-auditable": true}, "value": {"description": "Value.", "type": "string", "example": "alpn=\"h3,h2\" ipv4hint=\"127.0.0.1\" ipv6hint=\"::1\"", "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "SVCB", "enum": ["SVCB"], "x-auditable": true}}, "type": "object"}], "title": "SVCB Record"}
```
