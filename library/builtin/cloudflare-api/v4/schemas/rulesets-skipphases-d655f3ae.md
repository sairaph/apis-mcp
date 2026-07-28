---
title: rulesets_SkipPhases
page_id: schema-rulesets-skipphases-d655f3ae
path: schemas
description: A list of phases to skip the execution of. This option is incompatible with the rulesets option.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_SkipPhases

A list of phases to skip the execution of. This option is incompatible with the rulesets option.

```yaml
{"description": "A list of phases to skip the execution of. This option is incompatible with the rulesets option.", "type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/rulesets_RulesetPhase"}, {"description": "The phase to skip the execution of."}]}, "minItems": 1, "title": "Phases", "uniqueItems": true}
```
