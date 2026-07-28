---
title: dlp_EmailRuleCondition
page_id: schema-dlp-emailrulecondition-71fb1cdd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dlp_EmailRuleCondition

```yaml
{"type": "object", "properties": {"operator": {"$ref": "#/components/schemas/dlp_EmailRuleOperator"}, "selector": {"$ref": "#/components/schemas/dlp_EmailRuleSelector"}, "value": {"$ref": "#/components/schemas/dlp_EmailRuleValue"}}, "required": ["selector", "operator", "value"]}
```
