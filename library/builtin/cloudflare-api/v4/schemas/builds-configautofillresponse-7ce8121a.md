---
title: builds_ConfigAutofillResponse
page_id: schema-builds-configautofillresponse-7ce8121a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# builds_ConfigAutofillResponse

```yaml
{"type": "object", "properties": {"config_file": {"type": "string", "example": "wrangler.toml", "nullable": true}, "default_worker_name": {"type": "string", "example": "my-worker", "nullable": true}, "env_worker_names": {"type": "object", "example": {"production": "my-worker-prod", "staging": "my-worker-staging"}, "additionalProperties": {"type": "string"}, "nullable": true}, "package_manager": {"type": "string", "allOf": [{"$ref": "#/components/schemas/builds_PackageManager"}], "nullable": true}, "scripts": {"type": "object", "example": {"build": "npm run build", "test": "npm test"}, "additionalProperties": {"type": "string"}, "nullable": true}}}
```
