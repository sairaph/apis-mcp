---
title: rum_create-rule-request
page_id: schema-rum-create-rule-request-797ab3f2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rum_create-rule-request

```yaml
{"type": "object", "properties": {"host": {"type": "string", "example": "example.com"}, "inclusive": {"description": "Whether the rule includes or excludes traffic from being measured.", "type": "boolean", "example": true}, "is_paused": {"description": "Whether the rule is paused or not.", "type": "boolean", "example": false}, "paths": {"type": "array", "items": {"type": "string"}, "example": ["*"]}}}
```
