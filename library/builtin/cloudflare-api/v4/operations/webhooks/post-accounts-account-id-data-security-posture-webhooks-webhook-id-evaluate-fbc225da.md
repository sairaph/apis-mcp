---
title: Test an existing webhook configuration
page_id: operation-post-accounts-account-id-data-security-posture-webhooks-webhook-id-evalu-a3f28f54
path: operations/webhooks
description: "Sends a test webhook event using an existing webhook configuration.\nThis allows customers to verify their webhook endpoint is still reachable and properly \nconfigured after creating the webhook resource.\n\nThe test payload includes:\n- event_type: \"webhook.test\"\n- timestamp: Current UTC timestamp\n- message: Test message indicating this is from Cloudflare CASB\n- data: Object with test: true"
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}/evaluate
operation_ids:
    - EvaluateExistingWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Test an existing webhook configuration

`POST /accounts/{account_id}/data-security/posture/webhooks/{webhook_id}/evaluate`

Operation ID: `EvaluateExistingWebhook`

Sends a test webhook event using an existing webhook configuration.
This allows customers to verify their webhook endpoint is still reachable and properly
configured after creating the webhook resource.

The test payload includes:
- event_type: "webhook.test"
- timestamp: Current UTC timestamp
- message: Test message indicating this is from Cloudflare CASB
- data: Object with test: true

## Definition

```yaml
{"operationId": "EvaluateExistingWebhook", "summary": "Test an existing webhook configuration", "description": "Sends a test webhook event using an existing webhook configuration.\nThis allows customers to verify their webhook endpoint is still reachable and properly \nconfigured after creating the webhook resource.\n\nThe test payload includes:\n- event_type: \"webhook.test\"\n- timestamp: Current UTC timestamp\n- message: Test message indicating this is from Cloudflare CASB\n- data: Object with test: true\n", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}, {"$ref": "#/components/parameters/posture-api_WebhookId"}], "responses": {"200": {"description": "OK: Webhook test completed (check success field for result)", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/posture-api_EvaluateNewWebhookResponse"}}, "type": "object"}]}}}}, "400": {"description": "Bad Request: Invalid webhook ID", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "404": {"description": "Not Found: Webhook not found", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
