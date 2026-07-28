---
title: dns-records_OPENPGPKEYRecord
page_id: schema-dns-records-openpgpkeyrecord-12afbdfc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_OPENPGPKEYRecord

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns-record-shared-fields"}, {"properties": {"content": {"description": "A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)", "type": "string", "x-auditable": true}, "type": {"description": "Record type.", "type": "string", "example": "OPENPGPKEY", "enum": ["OPENPGPKEY"], "x-auditable": true}}, "type": "object"}], "title": "OPENPGPKEY Record"}
```
