---
title: rlanalytics_api-response-common-failure
page_id: schema-rlanalytics-api-response-common-failure-8ba92db1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rlanalytics_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/rlanalytics_messages"}, "minItems": 1}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/rlanalytics_messages"}}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages"]}
```
