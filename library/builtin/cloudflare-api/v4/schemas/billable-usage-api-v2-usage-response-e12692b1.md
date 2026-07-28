---
title: billable-usage-api_v2_usage_response
page_id: schema-billable-usage-api-v2-usage-response-e12692b1
path: schemas
description: Successful response containing an array of FOCUS-aligned cost and usage records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_v2_usage_response

Successful response containing an array of FOCUS-aligned cost and usage records.

```yaml
{"description": "Successful response containing an array of FOCUS-aligned cost and usage records.", "type": "object", "properties": {"errors": {"description": "Contains error details if the request failed.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "messages": {"description": "Contains informational notices about the response.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "result": {"description": "Contains the array of cost and usage records.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_cost_and_usage_data"}}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
