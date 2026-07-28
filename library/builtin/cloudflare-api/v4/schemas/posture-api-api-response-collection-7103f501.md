---
title: posture-api_api-response-collection
page_id: schema-posture-api-api-response-collection-7103f501
path: schemas
description: Response structure for paginated collections.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_api-response-collection

Response structure for paginated collections.

```yaml
{"description": "Response structure for paginated collections.", "type": "object", "properties": {"errors": {"$ref": "#/components/schemas/posture-api_messages"}, "messages": {"$ref": "#/components/schemas/posture-api_messages"}, "result_info": {"$ref": "#/components/schemas/posture-api_result-info"}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "errors", "messages", "result_info"]}
```
