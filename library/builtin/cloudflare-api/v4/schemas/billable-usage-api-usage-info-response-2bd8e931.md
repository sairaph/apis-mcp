---
title: billable-usage-api_usage_info_response
page_id: schema-billable-usage-api-usage-info-response-2bd8e931
path: schemas
description: Represents a successful response containing subscription info.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# billable-usage-api_usage_info_response

Represents a successful response containing subscription info.

```yaml
{"description": "Represents a successful response containing subscription info.", "type": "object", "properties": {"errors": {"description": "Contains error details if the request failed.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "messages": {"description": "Contains any informational messages from the API.", "type": "array", "items": {"$ref": "#/components/schemas/billable-usage-api_message"}, "nullable": true}, "result": {"description": "Contains the paygo usage info.", "type": "object", "properties": {"covered": {"description": "Indicates whether the account is covered.", "type": "boolean", "example": true}, "subscriptions": {"description": "List of subscriptions for the account.", "type": "array", "items": {"properties": {"billing_cycle_anchor_timestamp": {"description": "The subscription billing cycle anchor timestamp.", "type": "string", "format": "date-time", "example": "2023-01-01T00:00:00Z"}, "end_timestamp": {"description": "The subscription end timestamp. Omitted for active subscriptions; present only when the subscription has been cancelled.", "type": "string", "format": "date-time", "example": "2023-12-31T23:59:59Z"}, "id": {"description": "The identifier for the Cloudflare subscription.", "type": "string", "example": "3F3CD4CQ6N7FXO7IK6NVFJBOYA"}, "start_timestamp": {"description": "The subscription start timestamp.", "type": "string", "format": "date-time", "example": "2023-01-01T00:00:00Z"}}, "required": ["id", "start_timestamp", "billing_cycle_anchor_timestamp"], "type": "object"}}}, "required": ["covered", "subscriptions"]}, "success": {"description": "Indicates whether the API call was successful.", "type": "boolean", "example": true, "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}
```
