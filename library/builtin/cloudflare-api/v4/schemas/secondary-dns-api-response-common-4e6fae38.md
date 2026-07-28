---
title: secondary-dns_api-response-common
page_id: schema-secondary-dns-api-response-common-4e6fae38
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# secondary-dns_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/secondary-dns_messages"}, "messages": {"$ref": "#/components/schemas/secondary-dns_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
