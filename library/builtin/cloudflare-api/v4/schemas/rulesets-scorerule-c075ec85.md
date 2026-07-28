---
title: rulesets_ScoreRule
page_id: schema-rulesets-scorerule-c075ec85
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ScoreRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["score"]}, "action_parameters": {"properties": {"increment": {"$ref": "#/components/schemas/rulesets_ScoreIncrement"}}, "required": ["increment"]}, "description": {"example": "Increment the cumulative score."}}, "title": "Score Rule"}]}
```
