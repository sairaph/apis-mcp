---
title: workers_binding_kind_secret_key
page_id: schema-workers-binding-kind-secret-key-0d3aed0b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_secret_key

```yaml
{"type": "object", "properties": {"algorithm": {"description": "Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).", "type": "object", "x-auditable": true}, "format": {"description": "Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).", "type": "string", "example": "raw", "enum": ["raw", "pkcs8", "spki", "jwk"], "x-auditable": true}, "key_base64": {"description": "Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".", "type": "string", "writeOnly": true, "x-sensitive": true}, "key_jwk": {"description": "Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".", "type": "object", "writeOnly": true, "x-sensitive": true}, "name": {"$ref": "#/components/schemas/workers_binding_name"}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["secret_key"], "x-auditable": true}, "usages": {"description": "Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).", "type": "array", "items": {"enum": ["encrypt", "decrypt", "sign", "verify", "deriveKey", "deriveBits", "wrapKey", "unwrapKey"], "type": "string"}, "example": ["encrypt", "decrypt"], "x-auditable": true, "x-stainless-collection-type": "set"}}, "required": ["name", "type", "format", "algorithm", "usages"]}
```
