---
title: builds_EnvironmentVariablesRequest
page_id: schema-builds-environmentvariablesrequest-fb7a4ed7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_EnvironmentVariablesRequest

```yaml
{"type": "object", "example": {"API_KEY": {"is_secret": true, "value": "secret-key"}, "NODE_ENV": {"is_secret": false, "value": "production"}}, "additionalProperties": {"properties": {"is_secret": {"$ref": "#/components/schemas/builds_is_secret"}, "value": {"type": "string", "example": "production", "nullable": true}}, "required": ["is_secret"], "type": "object"}}
```
