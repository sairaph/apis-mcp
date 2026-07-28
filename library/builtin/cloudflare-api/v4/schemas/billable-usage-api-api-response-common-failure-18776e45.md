---
title: billable-usage-api_api-response-common-failure
page_id: schema-billable-usage-api-api-response-common-failure-18776e45
path: schemas
description: Represents a failed API response.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_api-response-common-failure

Represents a failed API response.

```yaml
{"description": "Represents a failed API response.", "type": "object", "properties": {"errors": {"description": "Contains error details describing why the request failed.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "example": [{"code": 1183, "message": "Invalid request parameters: Please ensure all required parameters are included and correctly formatted."}], "minItems": 1}, "messages": {"description": "Contains informational notices about the response.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "result": {"description": "Contains the response payload (always null on failure).", "type": "object", "nullable": true}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
