---
title: cache-purge_api-response-single-id
page_id: schema-cache-purge-api-response-single-id-bf62c9c3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-purge_api-response-single-id

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cache-purge_messages"}, "messages": {"$ref": "#/components/schemas/cache-purge_messages"}, "result": {"type": "object", "nullable": true, "properties": {"id": {"$ref": "#/components/schemas/cache-purge_identifier"}}, "required": ["id"]}, "success": {"description": "Indicates the API call's success or failure.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
