---
title: rulesets_ExecuteRule
page_id: schema-rulesets-executerule-2fafcdfc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ExecuteRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["execute"]}, "action_parameters": {"properties": {"id": {"allOf": [{"$ref": "#/components/schemas/rulesets_RulesetId"}, {"description": "The ID of the ruleset to execute.", "example": "4814384a9e5d4991b9815dcfc25d2f1f"}]}, "matched_data": {"$ref": "#/components/schemas/rulesets_ExecuteMatchedData"}, "overrides": {"$ref": "#/components/schemas/rulesets_ExecuteOverrides"}}, "required": ["id"]}, "description": {"example": "Execute another ruleset."}}, "title": "Execute Rule"}]}
```
