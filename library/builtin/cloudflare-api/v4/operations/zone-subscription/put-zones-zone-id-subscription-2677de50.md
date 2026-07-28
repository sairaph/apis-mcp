---
title: Update Zone Subscription
page_id: operation-put-zones-zone-id-subscription-06922d0e
path: operations/zone-subscription
description: Updates zone subscriptions, either plan or add-ons.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/subscription
operation_ids:
    - zone-subscription-update-zone-subscription
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zone Subscription

`PUT /zones/{zone_id}/subscription`

Operation ID: `zone-subscription-update-zone-subscription`

Updates zone subscriptions, either plan or add-ons.

## Definition

```yaml
{"operationId": "zone-subscription-update-zone-subscription", "summary": "Update Zone Subscription", "description": "Updates zone subscriptions, either plan or add-ons.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Update Zone Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}}}}, "4XX": {"description": "Update Zone Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.subscriptions", "x-fern-sdk-method-name": "update"}
```
