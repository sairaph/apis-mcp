---
title: dns-records_AAAARecord
page_id: schema-dns-records-aaaarecord-07ef7967
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_AAAARecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A valid IPv6 address.", "example": "2400:cb00:2049::1", "format": "ipv6", "type": "string", "x-auditable": true}, "private_routing": {"description": "Enables private network routing to the origin.", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "AAAA", "enum": ["AAAA"], "x-auditable": true}}, "type": "object"}], "title": "AAAA Record"}
```
