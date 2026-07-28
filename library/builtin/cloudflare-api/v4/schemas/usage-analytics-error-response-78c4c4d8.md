---
title: usage-analytics_error_response
page_id: schema-usage-analytics-error-response-78c4c4d8
path: schemas
description: Standard Cloudflare v4 API error response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# usage-analytics_error_response

Standard Cloudflare v4 API error response.

```yaml
{"description": "Standard Cloudflare v4 API error response.", "type": "object", "properties": {"errors": {"description": "List of errors.", "type": "array", "items": {"$ref": "#/components/schemas/usage-analytics_api_error"}}, "messages": {"description": "List of informational messages.", "type": "array", "items": {"$ref": "#/components/schemas/usage-analytics_api_message"}}, "result": {"description": "Empty array or null on error.", "type": "array", "items": {"type": "object"}, "example": [], "nullable": true}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": false}}, "required": ["success", "errors"]}
```
