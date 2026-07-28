---
title: MessagesErrorResponse
page_id: schema-messageserrorresponse-62925b1d
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesErrorResponse

```yaml
{"example": {"error": {"message": "Invalid request parameters", "type": "invalid_request_error"}, "type": "error"}, "properties": {"error": {"$ref": "#/components/schemas/MessagesErrorDetail"}, "type": {"enum": ["error"], "type": "string"}}, "required": ["type", "error"], "type": "object"}
```
