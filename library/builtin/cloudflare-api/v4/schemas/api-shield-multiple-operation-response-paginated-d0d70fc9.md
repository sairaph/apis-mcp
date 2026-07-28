---
title: api-shield_multiple-operation-response-paginated
page_id: schema-api-shield-multiple-operation-response-paginated-d0d70fc9
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_multiple-operation-response-paginated

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/api-shield_operation"}]}}}, "required": ["result"], "type": "object"}]}
```
