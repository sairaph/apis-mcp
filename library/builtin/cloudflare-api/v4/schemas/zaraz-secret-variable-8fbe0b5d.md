---
title: zaraz_secret_variable
page_id: schema-zaraz-secret-variable-8fbe0b5d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_secret_variable

```yaml
{"type": "object", "properties": {"name": {"type": "string", "x-auditable": true}, "type": {"type": "string", "enum": ["secret"], "x-auditable": true}, "value": {"type": "string", "x-auditable": true, "x-sensitive": true}}, "required": ["name", "type", "value"]}
```
