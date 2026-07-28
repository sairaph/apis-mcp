---
title: api-shield_credentials-JWT-Key-EC-common
page_id: schema-api-shield-credentials-jwt-key-ec-common-14157b5a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_credentials-JWT-Key-EC-common

```yaml
{"type": "object", "properties": {"kty": {"description": "Key Type", "type": "string", "enum": ["EC"], "x-auditable": true}, "x": {"description": "X EC coordinate", "type": "string", "x-auditable": false}, "y": {"description": "Y EC coordinate", "type": "string", "x-auditable": false}}, "required": ["kty", "x", "y"]}
```
