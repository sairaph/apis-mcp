---
title: Create a Notification policy
page_id: operation-post-accounts-account-id-alerting-v3-policies-3422a167
path: operations/notification-policies
description: Creates a new Notification policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies
operation_ids:
    - notification-policies-create-a-notification-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Notification policy

`POST /accounts/{account_id}/alerting/v3/policies`

Operation ID: `notification-policies-create-a-notification-policy`

Creates a new Notification policy.

## Definition

```yaml
{"operationId": "notification-policies-create-a-notification-policy", "summary": "Create a Notification policy", "description": "Creates a new Notification policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"alert_interval": {"$ref": "#/components/schemas/aaa_alert_interval"}, "alert_type": {"$ref": "#/components/schemas/aaa_alert_type"}, "description": {"$ref": "#/components/schemas/aaa_schemas-description"}, "enabled": {"$ref": "#/components/schemas/aaa_enabled"}, "filters": {"$ref": "#/components/schemas/aaa_filters"}, "mechanisms": {"$ref": "#/components/schemas/aaa_mechanisms"}, "name": {"$ref": "#/components/schemas/aaa_schemas-name"}}, "required": ["name", "alert_type", "enabled", "mechanisms"]}}}}, "responses": {"200": {"description": "Create a Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_id_response"}}}}, "4XX": {"description": "Create a Notification policy response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_id_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification policies"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
