---
title: workers-kv_api-response-common-failure
page_id: schema-workers-kv-api-response-common-failure-b1c83879
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/workers-kv_messages"}], "minLength": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/workers-kv_messages"}]}, "result": {"type": "object", "nullable": true, "x-stainless-empty-object": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
