---
title: List webhooks
page_id: operation-get-accounts-account-id-alerting-v3-destinations-webhooks-dd452397
path: operations/notification-webhooks
description: Gets a list of all configured webhook destinations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/destinations/webhooks
operation_ids:
    - notification-webhooks-list-webhooks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List webhooks

`GET /accounts/{account_id}/alerting/v3/destinations/webhooks`

Operation ID: `notification-webhooks-list-webhooks`

Gets a list of all configured webhook destinations.

## Definition

```yaml
{"operationId": "notification-webhooks-list-webhooks", "summary": "List webhooks", "description": "Gets a list of all configured webhook destinations.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "responses": {"200": {"description": "List webhooks response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_webhooks_components-schemas-response_collection"}}}}, "4XX": {"description": "List webhooks response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_webhooks_components-schemas-response_collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification webhooks"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
