---
title: csam-config-service_api_response_common_failure
page_id: schema-csam-config-service-api-response-common-failure-b1cc5e1f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# csam-config-service_api_response_common_failure

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/csam-config-service_api_response_message"}, "example": [{"code": 7003, "message": "No route for the URI"}]}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/csam-config-service_api_response_message"}}, "result": {"type": "object", "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false}}, "required": ["success", "errors", "messages"]}
```
