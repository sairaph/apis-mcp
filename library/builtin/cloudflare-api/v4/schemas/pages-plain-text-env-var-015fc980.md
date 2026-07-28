---
title: pages_plain_text_env_var
page_id: schema-pages-plain-text-env-var-015fc980
path: schemas
description: A plaintext environment variable.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# pages_plain_text_env_var

A plaintext environment variable.

```yaml
{"description": "A plaintext environment variable.", "type": "object", "properties": {"type": {"type": "string", "enum": ["plain_text"], "x-auditable": true}, "value": {"description": "Environment variable value.", "type": "string"}}, "example": {"type": "plain_text", "value": "hello world"}, "nullable": true, "required": ["type", "value"]}
```
