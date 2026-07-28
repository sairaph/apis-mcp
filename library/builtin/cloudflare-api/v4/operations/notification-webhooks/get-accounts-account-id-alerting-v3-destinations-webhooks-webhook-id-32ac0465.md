---
title: Get a webhook
page_id: operation-get-accounts-account-id-alerting-v3-destinations-webhooks-webhook-id-3e0f364c
path: operations/notification-webhooks
description: Get details for a single webhooks destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
operation_ids:
    - notification-webhooks-get-a-webhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a webhook

`GET /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}`

Operation ID: `notification-webhooks-get-a-webhook`

Get details for a single webhooks destination.

## Definition

```yaml
{"operationId": "notification-webhooks-get-a-webhook", "summary": "Get a webhook", "description": "Get details for a single webhooks destination.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "webhook_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_webhook-id"}}], "responses": {"200": {"description": "Get a webhook response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_schemas-single_response"}}}}, "4XX": {"description": "Get a webhook response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_schemas-single_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification webhooks"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
