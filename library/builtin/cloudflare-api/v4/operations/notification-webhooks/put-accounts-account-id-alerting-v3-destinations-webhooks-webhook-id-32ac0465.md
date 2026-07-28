---
title: Update a webhook
page_id: operation-put-accounts-account-id-alerting-v3-destinations-webhooks-webhook-id-f565fe54
path: operations/notification-webhooks
description: Update a webhook destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}
operation_ids:
    - notification-webhooks-update-a-webhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a webhook

`PUT /accounts/{account_id}/alerting/v3/destinations/webhooks/{webhook_id}`

Operation ID: `notification-webhooks-update-a-webhook`

Update a webhook destination.

## Definition

```yaml
{"operationId": "notification-webhooks-update-a-webhook", "summary": "Update a webhook", "description": "Update a webhook destination.", "parameters": [{"name": "webhook_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_webhook-id"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"name": {"$ref": "#/components/schemas/aaa_components-schemas-name"}, "secret": {"$ref": "#/components/schemas/aaa_secret"}, "url": {"$ref": "#/components/schemas/aaa_url"}}, "required": ["name", "url"]}}}}, "responses": {"200": {"description": "Update a webhook response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_id_response"}}}}, "4XX": {"description": "Update a webhook response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_id_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification webhooks"], "x-api-token-group": ["Notifications Write", "Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
