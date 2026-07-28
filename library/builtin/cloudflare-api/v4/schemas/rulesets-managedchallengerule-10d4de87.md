---
title: rulesets_ManagedChallengeRule
page_id: schema-rulesets-managedchallengerule-10d4de87
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_ManagedChallengeRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["managed_challenge"]}, "description": {"example": "Issue a Managed Challenge if the visitor has not solved a Managed Challenge or Interactive Challenge prior to the request."}}, "title": "Managed Challenge Rule"}]}
```
