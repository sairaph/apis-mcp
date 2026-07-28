---
title: cache_auto_origin_tls_kex_failure_response
page_id: schema-cache-auto-origin-tls-kex-failure-response-ac836f6d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache_auto_origin_tls_kex_failure_response

```yaml
{"type": "object", "properties": {"errors": {"allOf": [{"$ref": "#/components/schemas/cache_messages"}], "minLength": 1}, "messages": {"$ref": "#/components/schemas/cache_messages"}, "result": {"$ref": "#/components/schemas/cache_auto_origin_tls_kex_result"}, "success": {"description": "Indicates the API call's success or failure.", "type": "boolean", "example": false}}, "required": ["success", "errors", "messages", "result"]}
```
