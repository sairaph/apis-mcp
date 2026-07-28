---
title: rulesets_ExecuteOverrides
page_id: schema-rulesets-executeoverrides-b69f8b3b
path: schemas
description: A set of overrides to apply to the target ruleset.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ExecuteOverrides

A set of overrides to apply to the target ruleset.

```yaml
{"description": "A set of overrides to apply to the target ruleset.", "type": "object", "properties": {"action": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleAction"}, {"description": "An action to override all rules with. This option has lower precedence than rule and category overrides.", "example": "log"}]}, "categories": {"$ref": "#/components/schemas/rulesets_ExecuteCategoryOverrides"}, "enabled": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleEnabled"}, {"description": "Whether to enable execution of all rules. This option has lower precedence than rule and category overrides."}]}, "rules": {"$ref": "#/components/schemas/rulesets_ExecuteRuleOverrides"}, "sensitivity_level": {"allOf": [{"$ref": "#/components/schemas/rulesets_ExecuteSensitivityLevel"}, {"description": "A sensitivity level to set for all rules. This option has lower precedence than rule and category overrides and is only applicable for DDoS phases."}]}}, "minProperties": 1, "title": "Overrides"}
```
