---
title: cc_V4BaseResponse
page_id: schema-cc-v4baseresponse-0c5d98f2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_V4BaseResponse

```yaml
{"type": "object", "properties": {"errors": {"$ref": "#/components/schemas/cc_Messages"}, "messages": {"$ref": "#/components/schemas/cc_Messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
