---
title: api-shield_credentials-JWT-Key-oct-patch-request
page_id: schema-api-shield-credentials-jwt-key-oct-patch-request-7b41c697
path: schemas
description: JSON representation of a symmetric key for PATCH requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-oct-patch-request

JSON representation of a symmetric key for PATCH requests.

```yaml
{"description": "JSON representation of a symmetric key for PATCH requests.", "type": "object", "allOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-common"}, {"properties": {"alg": {"description": "Algorithm", "type": "string", "enum": ["HS256", "HS384", "HS512"], "x-auditable": true}, "k": {"description": "Symmetric key material. Optional for PATCH: omit to preserve existing secret for matching `{alg,kid}`; send a string to rotate. `k: null` is invalid.", "type": "string", "x-auditable": false}, "kty": {"description": "Key Type", "type": "string", "enum": ["oct"], "x-auditable": true}}, "required": ["kty", "alg"], "type": "object"}]}
```
