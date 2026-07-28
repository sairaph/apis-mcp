---
title: digital-experience-monitoring_dex_target_policy
page_id: schema-digital-experience-monitoring-dex-target-policy-7b4d172b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_dex_target_policy

```yaml
{"type": "object", "properties": {"default": {"description": "Whether the DEX rule is the account default.", "type": "boolean", "x-stainless-terraform-configurability": "computed"}, "id": {"description": "The id of the DEX rule.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}]}, "name": {"description": "The name of the DEX rule.", "type": "string", "x-stainless-terraform-configurability": "computed"}}, "required": ["id"]}
```
