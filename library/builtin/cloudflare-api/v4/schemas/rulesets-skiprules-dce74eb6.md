---
title: rulesets_SkipRules
page_id: schema-rulesets-skiprules-dce74eb6
path: schemas
description: A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of. This option is incompatible with the ruleset option.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SkipRules

A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of. This option is incompatible with the ruleset option.

```yaml
{"description": "A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the execution of. This option is incompatible with the ruleset option.", "type": "object", "example": {"4814384a9e5d4991b9815dcfc25d2f1f": ["8ac8bc2a661e475d940980f9317f28e1"]}, "additionalProperties": {"description": "A list of rule IDs in the ruleset to skip the execution of.", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleId"}, {"description": "The ID of a rule in the ruleset to skip the execution of.", "example": "8ac8bc2a661e475d940980f9317f28e1", "title": "Rule"}]}, "minItems": 1, "title": "Rules", "type": "array", "uniqueItems": true}, "minProperties": 1, "title": "Rules"}
```
