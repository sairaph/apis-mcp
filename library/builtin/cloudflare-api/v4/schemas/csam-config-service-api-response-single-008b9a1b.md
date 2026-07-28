---
title: csam-config-service_api_response_single
page_id: schema-csam-config-service-api-response-single-008b9a1b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_api_response_single

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/csam-config-service_api_response_message"}}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/csam-config-service_api_response_message"}}, "result": {"type": "object"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
