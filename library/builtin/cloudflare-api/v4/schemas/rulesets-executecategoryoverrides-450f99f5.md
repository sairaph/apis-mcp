---
title: rulesets_ExecuteCategoryOverrides
page_id: schema-rulesets-executecategoryoverrides-450f99f5
path: schemas
description: A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ExecuteCategoryOverrides

A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.

```yaml
{"description": "A list of category-level overrides. This option has the second-highest precedence after rule-level overrides.", "type": "array", "items": {"description": "A category-level override.", "minProperties": 2, "properties": {"action": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleAction"}, {"description": "The action to override rules in the category with.", "example": "log"}]}, "category": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleCategory"}, {"description": "The name of the category to override."}]}, "enabled": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleEnabled"}, {"description": "Whether to enable execution of rules in the category."}]}, "sensitivity_level": {"allOf": [{"$ref": "#/components/schemas/rulesets_ExecuteSensitivityLevel"}, {"description": "The sensitivity level to use for rules in the category. This option is only applicable for DDoS phases."}]}}, "required": ["category"], "title": "Category Override", "type": "object"}, "minItems": 1, "title": "Category Overrides", "uniqueItems": true}
```
