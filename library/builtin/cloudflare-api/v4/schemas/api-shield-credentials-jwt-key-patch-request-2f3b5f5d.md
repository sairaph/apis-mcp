---
title: api-shield_credentials-JWT-Key-patch-request
page_id: schema-api-shield-credentials-jwt-key-patch-request-2f3b5f5d
path: schemas
description: JSON representation of a JWKS key for PATCH credentials requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-patch-request

JSON representation of a JWKS key for PATCH credentials requests.

```yaml
{"description": "JSON representation of a JWKS key for PATCH credentials requests.", "type": "object", "oneOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-RSA"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-EC-ES256"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-EC-ES384"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-oct-patch-request"}]}
```
