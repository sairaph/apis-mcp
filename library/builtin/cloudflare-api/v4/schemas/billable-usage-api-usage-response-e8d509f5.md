---
title: billable-usage-api_usage_response
page_id: schema-billable-usage-api-usage-response-e8d509f5
path: schemas
description: Represents a successful response containing billable usage records.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_usage_response

Represents a successful response containing billable usage records.

```yaml
{"description": "Represents a successful response containing billable usage records.", "type": "object", "properties": {"errors": {"description": "Contains error details if the request failed.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "messages": {"description": "Contains informational notices about the response.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "result": {"description": "Contains the array of billable usage records.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_usage_record"}}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
