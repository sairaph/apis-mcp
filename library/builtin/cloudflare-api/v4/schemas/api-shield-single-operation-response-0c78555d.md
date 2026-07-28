---
title: api-shield_single-operation-response
page_id: schema-api-shield-single-operation-response-0c78555d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_single-operation-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/api-shield_api-response-common"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/api-shield_operation"}, {"properties": {"schemas": {"$ref": "#/components/schemas/api-shield_operation_schemas"}}, "type": "object"}]}}, "required": ["result"], "type": "object"}]}
```
