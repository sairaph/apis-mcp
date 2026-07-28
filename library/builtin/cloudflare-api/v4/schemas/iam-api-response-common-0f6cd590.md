---
title: iam_api-response-common
page_id: schema-iam-api-response-common-0f6cd590
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_api-response-common

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/iam_messages-2"}, "messages": {"$ref": "#/components/schemas/iam_messages-2"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages"]}
```
