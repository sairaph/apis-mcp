---
title: organizations-api_Entitlement
page_id: schema-organizations-api-entitlement-48e410d5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_Entitlement

```yaml
{"type": "object", "properties": {"allocation": {"anyOf": [{"$ref": "#/components/schemas/organizations-api_MaxCountAllocation"}, {"$ref": "#/components/schemas/organizations-api_BoolAllocation"}, {"$ref": "#/components/schemas/organizations-api_NullAllocation"}]}, "feature": {"$ref": "#/components/schemas/organizations-api_Feature"}}, "required": ["feature", "allocation"]}
```
