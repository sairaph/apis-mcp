---
title: rulesets_SkipRule
page_id: schema-rulesets-skiprule-17ec5ced
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SkipRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["skip"]}, "action_parameters": {"minProperties": 1, "properties": {"phase": {"$ref": "#/components/schemas/rulesets_SkipPhase"}, "phases": {"$ref": "#/components/schemas/rulesets_SkipPhases"}, "products": {"$ref": "#/components/schemas/rulesets_SkipProducts"}, "rules": {"$ref": "#/components/schemas/rulesets_SkipRules"}, "ruleset": {"$ref": "#/components/schemas/rulesets_SkipRuleset"}, "rulesets": {"$ref": "#/components/schemas/rulesets_SkipRulesets"}}}, "description": {"example": "Skip executing rulesets, rules, phases, and other products."}}, "title": "Skip Rule"}]}
```
