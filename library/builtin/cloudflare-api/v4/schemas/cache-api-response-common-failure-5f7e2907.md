---
title: cache_api-response-common-failure
page_id: schema-cache-api-response-common-failure-5f7e2907
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/cache_messages"}], "minLength": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/cache_messages"}]}, "result": {"$ref": "#/components/schemas/cache_result"}, "success": {"description": "Indicates the API call's success or failure.", "type": "boolean", "example": false}}, "required": ["success", "errors", "messages", "result"]}
```
