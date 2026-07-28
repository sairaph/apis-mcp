---
title: digital-experience-monitoring_create_rule_body
page_id: schema-digital-experience-monitoring-create-rule-body-ef57a4cc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_create_rule_body

```yaml
{"type": "object", "properties": {"description": {"type": "string", "x-auditable": true}, "match": {"description": "The wirefilter expression to match.", "type": "string"}, "name": {"description": "The name of the Rule.", "type": "string", "x-auditable": true}}, "required": ["name", "match"]}
```
