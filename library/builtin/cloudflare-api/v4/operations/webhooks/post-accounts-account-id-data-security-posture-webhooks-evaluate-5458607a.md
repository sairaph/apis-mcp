---
title: Test a webhook configuration before creating it
page_id: operation-post-accounts-account-id-data-security-posture-webhooks-evaluate-55e973aa
path: operations/webhooks
description: "Sends a test webhook event to the specified destination URL to verify the webhook endpoint \nis reachable and properly configured. This allows customers to validate their webhook \nconfiguration before creating the actual webhook resource.\n\nThe test payload includes:\n- event_type: \"webhook.test\"\n- timestamp: Current UTC timestamp\n- message: Test message indicating this is from Cloudflare CASB\n- data: Object with test: true"
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/data-security/posture/webhooks/evaluate
operation_ids:
    - EvaluateNewWebhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Test a webhook configuration before creating it

`POST /accounts/{account_id}/data-security/posture/webhooks/evaluate`

Operation ID: `EvaluateNewWebhook`

Sends a test webhook event to the specified destination URL to verify the webhook endpoint
is reachable and properly configured. This allows customers to validate their webhook
configuration before creating the actual webhook resource.

The test payload includes:
- event_type: "webhook.test"
- timestamp: Current UTC timestamp
- message: Test message indicating this is from Cloudflare CASB
- data: Object with test: true

## Definition

```yaml
{"operationId": "EvaluateNewWebhook", "summary": "Test a webhook configuration before creating it", "description": "Sends a test webhook event to the specified destination URL to verify the webhook endpoint \nis reachable and properly configured. This allows customers to validate their webhook \nconfiguration before creating the actual webhook resource.\n\nThe test payload includes:\n- event_type: \"webhook.test\"\n- timestamp: Current UTC timestamp\n- message: Test message indicating this is from Cloudflare CASB\n- data: Object with test: true\n", "parameters": [{"$ref": "#/components/parameters/posture-api_AccountTag"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_EvaluateNewWebhookRequest"}}}}, "responses": {"200": {"description": "OK: Webhook test completed (check success field for result)", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/posture-api_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/posture-api_EvaluateNewWebhookResponse"}}, "type": "object"}]}}}}, "400": {"description": "Bad Request: Invalid request parameters", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "401": {"description": "Unauthorized: Authentication required", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}, "403": {"description": "Forbidden: Insufficient permissions", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/posture-api_error-response"}}}}}, "security": [{"api_token": []}], "tags": ["webhooks"], "x-api-token-group": ["Zero Trust Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "data-security.posture", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
