---
title: api-shield_credentials-JWT-Key-EC-ES384
page_id: schema-api-shield-credentials-jwt-key-ec-es384-327a8ea2
path: schemas
description: JSON representation of an ES384 key
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-EC-ES384

JSON representation of an ES384 key

```yaml
{"description": "JSON representation of an ES384 key", "type": "object", "allOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-common"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-EC-common"}, {"properties": {"alg": {"description": "Algorithm", "type": "string", "enum": ["ES384"], "x-auditable": true}, "crv": {"description": "Curve", "type": "string", "enum": ["P-384"], "x-auditable": true}}, "required": ["alg", "crv"], "type": "object"}]}
```
