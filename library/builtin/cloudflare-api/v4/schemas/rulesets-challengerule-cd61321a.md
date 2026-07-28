---
title: rulesets_ChallengeRule
page_id: schema-rulesets-challengerule-cd61321a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ChallengeRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["challenge"]}, "description": {"example": "Issue an Interactive Challenge if the visitor has not solved an Interactive Challenge prior to the request."}}, "title": "Challenge Rule"}]}
```
