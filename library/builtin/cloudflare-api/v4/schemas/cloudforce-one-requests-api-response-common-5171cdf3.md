---
title: cloudforce-one-requests_api-response-common
page_id: schema-cloudforce-one-requests-api-response-common-5171cdf3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one-requests_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cloudforce-one-requests_messages"}, "messages": {"$ref": "#/components/schemas/cloudforce-one-requests_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
