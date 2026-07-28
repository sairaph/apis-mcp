---
title: realtimekit_ErrorResponse
page_id: schema-realtimekit-errorresponse-69773665
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_ErrorResponse

```yaml
{"type": "object", "properties": {"error": {"description": "Object containing details of the error that occurred", "type": "object", "properties": {"code": {"description": "Error code", "type": "number", "example": 404}, "message": {"description": "Error message", "type": "string", "example": "Error: resource not found"}}, "required": ["code", "message"]}, "success": {"description": "Whether the operation succeeded or not", "type": "boolean", "example": false}}, "required": ["success", "error"]}
```
