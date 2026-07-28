---
title: rulesets_DDoSDynamicRule
page_id: schema-rulesets-ddosdynamicrule-ab10a32a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_DDoSDynamicRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["ddos_dynamic"]}, "description": {"example": "Perform a specific action according to a set of internal guidelines defined by Cloudflare."}}, "title": "DDoS Dynamic Rule"}]}
```
