---
title: art_SummaryResponse
page_id: schema-art-summaryresponse-2bbbb7c6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_SummaryResponse

```yaml
{"type": "object", "properties": {"errors": {"type": "array", "items": {"$ref": "#/components/schemas/art_APIError"}, "example": []}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/art_APIMessage"}}, "result": {"$ref": "#/components/schemas/art_SummaryResult"}, "success": {"type": "boolean", "example": true}}, "required": ["success", "result", "errors", "messages"]}
```
