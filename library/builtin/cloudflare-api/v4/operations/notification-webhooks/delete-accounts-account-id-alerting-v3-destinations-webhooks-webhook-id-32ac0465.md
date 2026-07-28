---
title: Delete a webhook
page_id: operation-delete-accounts-account-id-alerting-v3-destinations-webhooks-webhook-id-e30eb565
path: operations/notification-webhooks
description: Delete a configured webhook destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
operation_ids:
    - notification-webhooks-delete-a-webhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a webhook

`DELETE /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}`

Operation ID: `notification-webhooks-delete-a-webhook`

Delete a configured webhook destination.

## Definition

```yaml
{"operationId": "notification-webhooks-delete-a-webhook", "summary": "Delete a webhook", "description": "Delete a configured webhook destination.", "parameters": [{"name": "webhook_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_webhook-id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "Delete a webhook response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-common-2"}}}}, "4XX": {"description": "Delete a webhook response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification webhooks"], "x-api-token-group": ["Notifications Write", "Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
