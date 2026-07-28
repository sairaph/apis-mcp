---
title: dns-records_DNSKEYRecord
page_id: schema-dns-records-dnskeyrecord-e94702a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_DNSKEYRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Formatted DNSKEY content. See 'data' to set DNSKEY properties.", "readOnly": true, "type": "string", "x-auditable": true}, "data": {"description": "Components of a DNSKEY record.", "type": "object", "properties": {"algorithm": {"description": "Algorithm.", "type": "number", "example": 5, "maximum": 255, "minimum": 0, "x-auditable": true}, "flags": {"description": "Flags.", "type": "number", "example": 1, "maximum": 65535, "minimum": 0, "x-auditable": true}, "protocol": {"description": "Protocol.", "type": "number", "example": 3, "maximum": 255, "minimum": 0, "x-auditable": true}, "public_key": {"description": "Public Key.", "type": "string", "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "DNSKEY", "enum": ["DNSKEY"], "x-auditable": true}}, "type": "object"}], "title": "DNSKEY Record"}
```
