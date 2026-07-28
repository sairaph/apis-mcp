---
title: page-shield_api-response-common-failure
page_id: schema-page-shield-api-response-common-failure-7955b490
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# page-shield_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/page-shield_messages"}, "messages": {"$ref": "#/components/schemas/page-shield_messages"}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors"]}
```
