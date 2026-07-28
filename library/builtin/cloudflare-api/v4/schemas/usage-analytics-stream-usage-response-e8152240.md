---
title: usage-analytics_stream_usage_response
page_id: schema-usage-analytics-stream-usage-response-e8152240
path: schemas
description: Standard Cloudflare v4 API response containing Stream usage analytics data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# usage-analytics_stream_usage_response

Standard Cloudflare v4 API response containing Stream usage analytics data.

```yaml
{"description": "Standard Cloudflare v4 API response containing Stream usage analytics data.", "type": "object", "properties": {"errors": {"description": "List of errors, empty on success.", "type": "array", "items": {"$ref": "#/components/schemas/usage-analytics_api_error"}}, "messages": {"description": "List of informational messages.", "type": "array", "items": {"$ref": "#/components/schemas/usage-analytics_api_message"}}, "result": {"description": "Array of usage data points for the requested time range.", "type": "array", "items": {"$ref": "#/components/schemas/usage-analytics_stream_usage_data_point"}}, "success": {"description": "Whether the API call was successful.", "type": "boolean", "example": true}}, "required": ["success", "result", "errors"]}
```
