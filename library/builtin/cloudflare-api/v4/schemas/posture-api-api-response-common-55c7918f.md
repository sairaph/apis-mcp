---
title: posture-api_api-response-common
page_id: schema-posture-api-api-response-common-55c7918f
path: schemas
description: Common response structure for all API endpoints.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_api-response-common

Common response structure for all API endpoints.

```yaml
{"description": "Common response structure for all API endpoints.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/posture-api_messages"}, "messages": {"$ref": "#/components/schemas/posture-api_messages"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages"]}
```
