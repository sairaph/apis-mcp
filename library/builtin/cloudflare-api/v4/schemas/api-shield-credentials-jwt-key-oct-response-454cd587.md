---
title: api-shield_credentials-JWT-Key-oct-response
page_id: schema-api-shield-credentials-jwt-key-oct-response-454cd587
path: schemas
description: JSON representation of a symmetric verification key in API responses (secret material is redacted).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-oct-response

JSON representation of a symmetric verification key in API responses (secret material is redacted).

```yaml
{"description": "JSON representation of a symmetric verification key in API responses (secret material is redacted).", "type": "object", "allOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-common"}, {"properties": {"alg": {"description": "Algorithm", "type": "string", "enum": ["HS256", "HS384", "HS512"], "x-auditable": true}, "kty": {"description": "Key Type", "type": "string", "enum": ["oct"], "x-auditable": true}}, "required": ["kty", "alg"], "type": "object"}]}
```
