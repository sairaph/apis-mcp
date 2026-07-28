---
title: dnssec_dnskey_record
page_id: schema-dnssec-dnskey-record-26df85e7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnskey_record

```yaml
{"type": "object", "properties": {"Algorithm": {"type": "integer", "example": 13, "nullable": true, "readOnly": true}, "Flags": {"type": "integer", "example": 256, "nullable": true, "readOnly": true}, "Hdr": {"$ref": "#/components/schemas/dnssec_dnskey_record_header"}, "Protocol": {"type": "integer", "example": 3, "nullable": true, "readOnly": true}, "PublicKey": {"type": "string", "example": "oXiGYrSTO+LSCJ3mohc8EP+CzF9KxBj8/ydXJ22pKuZP3VAC3/Md/k7xZfz470CoRyZJ6gV6vml07IC3d8xqhA==", "nullable": true, "readOnly": true}}}
```
