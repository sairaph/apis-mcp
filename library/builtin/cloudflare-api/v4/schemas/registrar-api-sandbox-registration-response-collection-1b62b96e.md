---
title: registrar-api-sandbox_registration-response-collection
page_id: schema-registrar-api-sandbox-registration-response-collection-1b62b96e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_registration-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/registrar-api-sandbox_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/registrar-api-sandbox_registration"}}, "result_info": {"$ref": "#/components/schemas/registrar-api-sandbox_cursor_result_info"}}, "required": ["result", "result_info"], "type": "object"}]}
```
