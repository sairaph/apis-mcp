---
title: rulesets_ExecuteRuleOverrides
page_id: schema-rulesets-executeruleoverrides-e40fc042
path: schemas
description: A list of rule-level overrides. This option has the highest precedence.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ExecuteRuleOverrides

A list of rule-level overrides. This option has the highest precedence.

```yaml
{"description": "A list of rule-level overrides. This option has the highest precedence.", "type": "array", "items": {"description": "A rule-level override.", "minProperties": 2, "properties": {"action": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleAction"}, {"description": "The action to override the rule with.", "example": "log"}]}, "enabled": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleEnabled"}, {"description": "Whether to enable execution of the rule."}]}, "id": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleId"}, {"description": "The ID of the rule to override.", "example": "8ac8bc2a661e475d940980f9317f28e1"}]}, "score_threshold": {"description": "The score threshold to use for the rule.", "type": "integer", "title": "Score Threshold"}, "sensitivity_level": {"allOf": [{"$ref": "#/components/schemas/rulesets_ExecuteSensitivityLevel"}, {"description": "The sensitivity level to use for the rule. This option is only applicable for DDoS phases."}]}}, "required": ["id"], "title": "Rule Override", "type": "object"}, "minItems": 1, "title": "Rule Overrides", "uniqueItems": true}
```
