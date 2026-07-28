---
title: zaraz_variable-match-rule
page_id: schema-zaraz-variable-match-rule-1416d12b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_variable-match-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["variableMatch"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"match": {"type": "string", "x-auditable": true}, "variable": {"type": "string", "x-auditable": true}}, "required": ["variable", "match"]}}, "required": ["id", "action", "settings"]}
```
