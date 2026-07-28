---
title: Zone Subscription Details
page_id: operation-get-zones-zone-id-subscription-77940622
path: operations/zone-subscription
description: Lists zone subscription details.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/subscription
operation_ids:
    - zone-subscription-zone-subscription-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Zone Subscription Details

`GET /zones/{zone_id}/subscription`

Operation ID: `zone-subscription-zone-subscription-details`

Lists zone subscription details.

## Definition

```yaml
{"operationId": "zone-subscription-zone-subscription-details", "summary": "Zone Subscription Details", "description": "Lists zone subscription details.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "responses": {"200": {"description": "Zone Subscription Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}}}}, "4XX": {"description": "Zone Subscription Details response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Subscription"], "x-api-token-group": ["Billing Write", "Billing Read"], "x-cfPermissionsRequired": {"enum": ["#billing:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.subscriptions", "x-fern-sdk-method-name": "get"}
```
