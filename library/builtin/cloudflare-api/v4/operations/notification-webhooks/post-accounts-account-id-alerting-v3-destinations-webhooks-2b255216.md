---
title: Create a webhook
page_id: operation-post-accounts-account-id-alerting-v3-destinations-webhooks-d455be1d
path: operations/notification-webhooks
description: Creates a new webhook destination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/webhooks
operation_ids:
    - notification-webhooks-create-a-webhook
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a webhook

`POST /accounts/{account_id}/alerting/v3/destinations/webhooks`

Operation ID: `notification-webhooks-create-a-webhook`

Creates a new webhook destination.

## Definition

```yaml
{"operationId": "notification-webhooks-create-a-webhook", "summary": "Create a webhook", "description": "Creates a new webhook destination.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"name": {"$ref": "#/components/schemas/aaa_components-schemas-name"}, "secret": {"$ref": "#/components/schemas/aaa_secret"}, "url": {"$ref": "#/components/schemas/aaa_url"}}, "required": ["name", "url"]}}}}, "responses": {"201": {"description": "Create a webhook response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_id_response"}}}}, "4XX": {"description": "Create a webhook response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_id_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification webhooks"], "x-api-token-group": ["Notifications Write", "Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
