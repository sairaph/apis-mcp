---
title: zaraz_element-visibility-rule
page_id: schema-zaraz-element-visibility-rule-610db464
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_element-visibility-rule

```yaml
{"type": "object", "properties": {"action": {"type": "string", "enum": ["elementVisibility"], "x-auditable": true}, "id": {"type": "string", "x-auditable": true}, "settings": {"type": "object", "properties": {"selector": {"type": "string", "x-auditable": true}}, "required": ["selector"]}}, "required": ["id", "action", "settings"]}
```
