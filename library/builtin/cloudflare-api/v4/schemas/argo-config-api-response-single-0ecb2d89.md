---
title: argo-config_api_response_single
page_id: schema-argo-config-api-response-single-0ecb2d89
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# argo-config_api_response_single

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/argo-config_messages"}, "messages": {"$ref": "#/components/schemas/argo-config_messages"}, "result": {"$ref": "#/components/schemas/argo-config_result_object"}, "success": {"description": "Describes a successful API response.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
