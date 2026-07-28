---
title: cache-purge_api-response-common-failure
page_id: schema-cache-purge-api-response-common-failure-c2f548c2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_api-response-common-failure

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/cache-purge_messages"}], "minLength": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/cache-purge_messages"}]}, "result": {"type": "object", "nullable": true}, "success": {"description": "Indicates the API call's success or failure.", "type": "boolean", "example": false}}, "required": ["success", "errors", "messages", "result"]}
```
