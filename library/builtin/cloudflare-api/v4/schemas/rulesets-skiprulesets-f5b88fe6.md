---
title: rulesets_SkipRulesets
page_id: schema-rulesets-skiprulesets-f5b88fe6
path: schemas
description: A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SkipRulesets

A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.

```yaml
{"description": "A list of ruleset IDs to skip the execution of. This option is incompatible with the ruleset and phases options.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_RulesetId"}, {"description": "The ID of a ruleset to skip the execution of.", "example": "4814384a9e5d4991b9815dcfc25d2f1f", "title": "Ruleset"}]}, "minItems": 1, "title": "Rulesets", "uniqueItems": true}
```
