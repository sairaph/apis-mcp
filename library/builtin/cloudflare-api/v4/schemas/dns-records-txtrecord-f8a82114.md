---
title: dns-records_TXTRecord
page_id: schema-dns-records-txtrecord-f8a82114
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_TXTRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "Text content for the record. The content must consist of quoted \"character strings\" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding this allowed maximum length are automatically split.\n\nLearn more at <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.", "example": "\"v=spf1 include:example.com -all\"", "type": "string", "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "TXT", "enum": ["TXT"], "x-auditable": true}}, "type": "object"}], "title": "TXT Record"}
```
