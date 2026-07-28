---
title: builds_EnvironmentVariablesResponse
page_id: schema-builds-environmentvariablesresponse-4525b325
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_EnvironmentVariablesResponse

```yaml
{"type": "object", "example": {"API_KEY": {"created_on": "2023-01-01T00:00:00Z", "is_secret": true, "value": null}, "NODE_ENV": {"created_on": "2023-01-01T00:00:00Z", "is_secret": false, "value": "production"}}, "additionalProperties": {"properties": {"created_on": {"$ref": "#/components/schemas/builds_created_on"}, "is_secret": {"$ref": "#/components/schemas/builds_is_secret"}, "value": {"description": "Value is null for secret environment variables", "type": "string", "nullable": true}}, "required": ["is_secret", "created_on"], "type": "object"}}
```
