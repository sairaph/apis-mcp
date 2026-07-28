---
title: api-shield_credentials-JWT-Key-response
page_id: schema-api-shield-credentials-jwt-key-response-2d003ee0
path: schemas
description: JSON representation of a JWKS key.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-response

JSON representation of a JWKS key.

```yaml
{"description": "JSON representation of a JWKS key.", "type": "object", "example": {"alg": "ES256", "crv": "P-256", "kid": "38013f13-c266-4eec-a72a-92ec92779f21", "kty": "EC", "x": "KN53JRwN3wCjm2o39bvZUX2VdrsHzS8pxOAGjm8m7EQ", "y": "lnkkzIxaveggz-HFhcMWW15nxvOj0Z_uQsXbpK0GFcY"}, "oneOf": [{"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-RSA"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-EC-ES256"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-EC-ES384"}, {"$ref": "#/components/schemas/api-shield_credentials-JWT-Key-oct-response"}]}
```
