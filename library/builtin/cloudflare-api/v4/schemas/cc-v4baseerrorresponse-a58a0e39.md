---
title: cc_V4BaseErrorResponse
page_id: schema-cc-v4baseerrorresponse-a58a0e39
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_V4BaseErrorResponse

```yaml
{"type": "object", "properties": {"errors": {"example": [{"code": 7003, "message": "No route for the URI"}], "allOf": [{"$ref": "#/components/schemas/cc_Messages"}], "minLength": 1}, "messages": {"example": [], "allOf": [{"$ref": "#/components/schemas/cc_Messages"}]}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean"}}, "required": ["success", "errors", "messages", "result"]}
```
