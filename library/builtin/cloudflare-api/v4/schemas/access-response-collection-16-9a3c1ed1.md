---
title: access_response_collection-16
page_id: schema-access-response-collection-16-9a3c1ed1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_response_collection-16

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/access_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/access_authenticator_device_aaguid"}, "example": [{"aaguid": "2fc0579f-8113-47ea-b116-bb5a8db9202a", "name": "YubiKey 5 NFC"}, {"aaguid": "fcb1bcb4-f370-078c-6993-bc24d0ae3fbe", "name": "Ledger Nano X FIDO2 Authenticator"}]}}, "type": "object"}]}
```
