---
title: dns-records_CNAMERecord
page_id: schema-dns-records-cnamerecord-ab3bf7e9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_CNAMERecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A valid hostname. Must not match the record's name.", "type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"flatten_cname": {"description": "If enabled, causes the CNAME record to be resolved externally and the resulting address records (e.g., A and AAAA) to be returned instead of the CNAME record itself. This setting is unavailable for proxied records, since they are always flattened.", "type": "boolean", "example": true, "default": false, "x-auditable": true}}}, "type": {"description": "Record type.", "type": "string", "example": "CNAME", "enum": ["CNAME"], "x-auditable": true}}, "type": "object"}], "title": "CNAME Record"}
```
