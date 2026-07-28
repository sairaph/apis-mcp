---
title: rulesets_JsChallengeRule
page_id: schema-rulesets-jschallengerule-df7deafb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_JsChallengeRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["js_challenge"]}, "description": {"example": "Issue a non-interactive JavaScript Challenge if the visitor has not solved an Interactive Challenge, Managed Challenge, or JavaScript Challenge prior to the request."}}, "title": "JavaScript Challenge Rule"}]}
```
