---
title: page-shield_policy-with-id
page_id: schema-page-shield-policy-with-id-4a28bf78
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_policy-with-id

```yaml
{"allOf": [{"$ref": "#/components/schemas/page-shield_policy"}, {"properties": {"id": {"$ref": "#/components/schemas/page-shield_id"}}, "required": ["id"]}], "required": ["id", "description", "action", "expression", "enabled", "value"]}
```
