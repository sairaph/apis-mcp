---
title: dnssec_dnssec_zsk
page_id: schema-dnssec-dnssec-zsk-0cb08fee
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnssec_zsk

```yaml
{"type": "object", "properties": {"DNSKEY": {"$ref": "#/components/schemas/dnssec_dnskey_record"}, "Location": {"$ref": "#/components/schemas/dnssec_key_storage_location"}, "Name": {"description": "Internal key name for the ZSK.", "type": "string", "example": "zsk_default", "readOnly": true}, "SigningKey": {"$ref": "#/components/schemas/dnssec_dnskey_signing_key"}, "Tag": {"$ref": "#/components/schemas/dnssec_dnssec_key_state"}}}
```
