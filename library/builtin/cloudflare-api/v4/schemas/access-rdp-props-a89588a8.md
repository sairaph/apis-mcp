---
title: access_rdp_props
page_id: schema-access-rdp-props-a89588a8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_rdp_props

```yaml
{"allOf": [{"$ref": "#/components/schemas/access_app_resp_embedded_target_criteria_self_hosted"}, {"$ref": "#/components/schemas/access_self_hosted_props"}, {"properties": {"type": {"allOf": [{"$ref": "#/components/schemas/access_type"}, {"example": "rdp"}]}}}], "required": ["target_criteria"]}
```
