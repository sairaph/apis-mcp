---
title: tls-certificates-and-hostnames_custom_hostname_api_response_common
page_id: schema-tls-certificates-and-hostnames-custom-hostname-api-response-common-6d7d7e40
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_hostname_api_response_common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_messages"}, "messages": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_hostname_response_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
