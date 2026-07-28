---
title: api-shield_auth_id_characteristic
page_id: schema-api-shield-auth-id-characteristic-c82c849b
path: schemas
description: Auth ID Characteristic
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_auth_id_characteristic

Auth ID Characteristic

```yaml
{"description": "Auth ID Characteristic", "properties": {"name": {"description": "The name of the characteristic field, i.e., the header or cookie name.", "type": "string", "example": "authorization", "maxLength": 128, "x-auditable": true}, "type": {"description": "The type of characteristic.", "type": "string", "example": "header", "enum": ["header", "cookie"], "x-auditable": true}}, "required": ["type", "name"]}
```
