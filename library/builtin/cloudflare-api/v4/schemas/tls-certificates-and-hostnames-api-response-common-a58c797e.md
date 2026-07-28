---
title: tls-certificates-and-hostnames_api-response-common
page_id: schema-tls-certificates-and-hostnames-api-response-common-a58c797e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_messages"}, "messages": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
