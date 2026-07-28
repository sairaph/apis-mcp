---
title: api-shield_credentials-JWT-Key-RSA
page_id: schema-api-shield-credentials-jwt-key-rsa-4d851d4d
path: schemas
description: JSON representation of an RSA key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-RSA

JSON representation of an RSA key.

```yaml
{"description": "JSON representation of an RSA key.", "type": "object", "allOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-common"}, {"properties": {"alg": {"description": "Algorithm", "type": "string", "enum": ["RS256", "RS384", "RS512", "PS256", "PS384", "PS512"], "x-auditable": true}, "e": {"description": "RSA exponent", "type": "string", "x-auditable": false}, "kty": {"description": "Key Type", "type": "string", "enum": ["RSA"], "x-auditable": true}, "n": {"description": "RSA modulus", "type": "string", "x-auditable": false}}, "required": ["kty", "alg", "n", "e"], "type": "object"}]}
```
