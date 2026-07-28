---
title: api-shield_auth_id_characteristic_jwt_claim
page_id: schema-api-shield-auth-id-characteristic-jwt-claim-d521b832
path: schemas
description: Auth ID Characteristic extracted from JWT Token Claims
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_auth_id_characteristic_jwt_claim

Auth ID Characteristic extracted from JWT Token Claims

```yaml
{"description": "Auth ID Characteristic extracted from JWT Token Claims", "type": "object", "properties": {"name": {"description": "Claim location expressed as `$(token_config_id):$(json_path)`, where `token_config_id` \nis the ID of the token configuration used in validating the JWT, and `json_path` is a RFC 9535 \nJSONPath (https://goessner.net/articles/JsonPath/, https://www.rfc-editor.org/rfc/rfc9535.html).\nThe JSONPath expression may be in dot or bracket notation, may only specify literal keys\nor array indexes, and must return a singleton value, which will be interpreted as a string.\n", "type": "string", "example": "e0de1a3a-8c2c-4f90-98d8-cbdf0a3f2cb5:$.foo.bar[0].baz", "maxLength": 128, "pattern": "^(?<token_config_id>[a-z0-9\\-]{32,36}):\\$(?<json_path>.*?)$", "x-auditable": true}, "type": {"description": "The type of characteristic.", "type": "string", "example": "jwt", "enum": ["jwt"], "x-auditable": true}}, "required": ["type", "name"]}
```
