---
title: dns-records_ARecord
page_id: schema-dns-records-arecord-1ed0bc40
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_ARecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A valid IPv4 address.", "example": "198.51.100.4", "format": "ipv4", "type": "string", "x-auditable": true}, "private_routing": {"description": "Enables private network routing to the origin.", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "A", "enum": ["A"], "x-auditable": true}}, "type": "object"}], "title": "A Record"}
```
