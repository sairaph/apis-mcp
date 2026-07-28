---
title: zaraz_zaraz-history-response
page_id: schema-zaraz-zaraz-history-response-071556fa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zaraz_zaraz-history-response

```yaml
{"allOf": [{"$ref": "#/components/schemas/zaraz_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"allOf": [{"$ref": "#/components/schemas/zaraz_zaraz-config-row-base"}, {"example": {"createdAt": "2023-02-23T05:05:55.155273Z", "description": "Config with enabled ecommerce tracking", "id": 12345, "updatedAt": "2023-02-23T05:05:55.155273Z", "userId": "278d0d0g123cd8e49d45ea64f12faa37"}, "properties": {"description": {"description": "Configuration description provided by the user who published this configuration.", "type": "string"}}, "required": ["description"], "type": "object"}]}}}, "type": "object"}]}
```
