---
title: rulesets_RedirectRule
page_id: schema-rulesets-redirectrule-2aa97937
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_RedirectRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["redirect"]}, "action_parameters": {"maxProperties": 1, "minProperties": 1, "properties": {"from_list": {"$ref": "#/components/schemas/rulesets_RedirectFromList"}, "from_value": {"$ref": "#/components/schemas/rulesets_RedirectFromValue"}}}, "description": {"example": "Redirect the request."}}, "title": "Redirect Rule"}]}
```
