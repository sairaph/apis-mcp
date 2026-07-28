---
title: rulesets_RewriteRule
page_id: schema-rulesets-rewriterule-d25668bd
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RewriteRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["rewrite"]}, "action_parameters": {"minProperties": 1, "properties": {"headers": {"$ref": "#/components/schemas/rulesets_RewriteHeaders"}, "uri": {"$ref": "#/components/schemas/rulesets_RewriteUri"}}}, "description": {"example": "Rewrite properties of the request or response."}}, "title": "Rewrite Rule"}]}
```
