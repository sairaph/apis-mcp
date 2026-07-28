---
title: posture-api_error-response
page_id: schema-posture-api-error-response-8ad76275
path: schemas
description: Standard error response structure.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_error-response

Standard error response structure.

```yaml
{"description": "Standard error response structure.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/posture-api_messages"}, "messages": {"$ref": "#/components/schemas/posture-api_messages"}, "success": {"description": "Indicates the request failed.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages"]}
```
