---
title: dnssec_dnskey_signing_key
page_id: schema-dnssec-dnskey-signing-key-266acfa8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dnssec_dnskey_signing_key

```yaml
{"type": "object", "properties": {"kek": {"description": "Key encryption key name used to encrypt the private key.", "type": "string", "example": "edge_kek_default", "nullable": true, "readOnly": true}, "privkey": {"description": "Encrypted private key material for the signing key.", "type": "string", "format": "byte", "example": "U3ZlbidzIHZlcnkgc2VjcmV0IGtleQ==", "nullable": true, "readOnly": true}, "pubkey": {"description": "Public key content associated with the signing key.", "type": "string", "example": "256 3 13 oXiGYrSTO+LSCJ3mohc8EP+CzF9KxBj8/ydXJ22pKuZP3VAC3/Md/k7xZfz470CoRyZJ6gV6vml07IC3d8xqhA==", "nullable": true, "readOnly": true}}}
```
