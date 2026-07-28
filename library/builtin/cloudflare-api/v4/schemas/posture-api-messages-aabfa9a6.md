---
title: posture-api_messages
page_id: schema-posture-api-messages-aabfa9a6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_messages

```yaml
{"type": "array", "items": {"properties": {"code": {"description": "Error or message code.", "type": "integer", "example": 1000, "minimum": 1000}, "documentation_url": {"description": "Link to relevant documentation.", "type": "string", "format": "uri", "example": "https://developers.cloudflare.com/api/operations/list-findings"}, "message": {"description": "Human-readable message.", "type": "string", "example": "Request processed successfully"}, "source": {"type": "object", "properties": {"pointer": {"description": "JSON pointer to the source of the error.", "type": "string", "example": "/data/attributes/name"}}}}, "required": ["code", "message"], "type": "object"}, "example": [], "uniqueItems": true}
```
