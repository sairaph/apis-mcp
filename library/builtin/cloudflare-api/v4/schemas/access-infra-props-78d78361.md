---
title: access_infra_props
page_id: schema-access-infra-props-78d78361
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_infra_props

```yaml
{"type": "object", "allOf": [{"properties": {"name": {"$ref": "#/components/schemas/access_name-8"}, "type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "infrastructure"}]}}}, {"$ref": "#/components/schemas/access_app_resp_embedded_target_criteria_infra"}], "required": ["type", "target_criteria"], "title": "Infrastructure Application"}
```
