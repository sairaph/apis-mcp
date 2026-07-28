---
title: api-shield_multiple-operation-response
page_id: schema-api-shield-multiple-operation-response-7d70de7c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_multiple-operation-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/api-shield_operation"}}}, "required": ["result"], "type": "object"}]}
```
