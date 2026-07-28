---
title: realtimekit_GenericErrorResponse
page_id: schema-realtimekit-genericerrorresponse-5a4d0be5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_GenericErrorResponse

```yaml
{"type": "object", "properties": {"error": {"type": "object", "properties": {"code": {"description": "HTTP status code of the error.", "type": "number"}, "message": {"description": "Error message describing what went wrong.", "type": "string"}}, "required": ["code", "message"]}, "success": {"description": "Success status of the request.", "type": "boolean", "default": false}}, "required": ["success", "error"]}
```
