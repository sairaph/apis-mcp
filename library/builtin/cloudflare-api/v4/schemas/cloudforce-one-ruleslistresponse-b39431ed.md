---
title: cloudforce-one_RulesListResponse
page_id: schema-cloudforce-one-ruleslistresponse-b39431ed
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_RulesListResponse

```yaml
{"type": "object", "properties": {"rules": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one_Rule"}}, "total": {"type": "number", "example": 100}}, "required": ["rules", "total"]}
```
