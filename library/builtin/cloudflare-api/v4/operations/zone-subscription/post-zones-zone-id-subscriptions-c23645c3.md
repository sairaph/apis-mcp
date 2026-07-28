---
title: Create Zone Subscription
page_id: operation-post-zones-zone-id-subscriptions-b6438ebb
path: operations/zone-subscription
description: Create a zone subscription, either plan or add-ons. Retained for audit-log coverage. Use the singular `/zones/{zone_id}/subscription` path instead.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/subscriptions
operation_ids:
    - zone-subscription-create-zone-subscriptions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zone Subscription

`POST /zones/{zone_id}/subscriptions`

Operation ID: `zone-subscription-create-zone-subscriptions`

Create a zone subscription, either plan or add-ons. Retained for audit-log coverage. Use the singular `/zones/{zone_id}/subscription` path instead.

## Definition

```yaml
{"operationId": "zone-subscription-create-zone-subscriptions", "summary": "Create Zone Subscription", "description": "Create a zone subscription, either plan or add-ons. Retained for audit-log coverage. Use the singular `/zones/{zone_id}/subscription` path instead.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/bill-subs-api_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_subscription-v2"}}}}, "responses": {"200": {"description": "Create Zone Subscription response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}}}}, "4XX": {"description": "Create Zone Subscription response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/bill-subs-api_zone_subscription_response_single"}, {"$ref": "#/components/schemas/bill-subs-api_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zone Subscription"], "x-api-token-group": ["Billing Write"], "x-cfPermissionsRequired": {"enum": ["#billing:read", "#billing:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-ignore": true, "x-forge-hidden": true}
```
