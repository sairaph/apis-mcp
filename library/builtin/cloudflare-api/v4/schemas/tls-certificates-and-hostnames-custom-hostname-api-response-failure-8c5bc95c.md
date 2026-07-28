---
title: tls-certificates-and-hostnames_custom_hostname_api_response_failure
page_id: schema-tls-certificates-and-hostnames-custom-hostname-api-response-failure-8c5bc95c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_hostname_api_response_failure

```yaml
{"type": "object", "properties": {"errors": {"allOf": [{"$ref": "#/components/schemas/tls-certificates-and-hostnames_messages"}, {"example": [{"code": 7003, "message": "No route for the URI"}], "minItems": 1}]}, "messages": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_messages"}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
