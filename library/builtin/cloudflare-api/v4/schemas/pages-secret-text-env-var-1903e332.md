---
title: pages_secret_text_env_var
page_id: schema-pages-secret-text-env-var-1903e332
path: schemas
description: An encrypted environment variable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_secret_text_env_var

An encrypted environment variable.

```yaml
{"description": "An encrypted environment variable.", "type": "object", "properties": {"type": {"type": "string", "enum": ["secret_text"], "x-auditable": true}, "value": {"description": "Secret value.", "type": "string", "x-sensitive": true}}, "example": {"type": "secret_text", "value": ""}, "nullable": true, "required": ["type", "value"]}
```
