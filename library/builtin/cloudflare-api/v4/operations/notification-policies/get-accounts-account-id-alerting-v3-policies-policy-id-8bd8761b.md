---
title: Get a Notification policy
page_id: operation-get-accounts-account-id-alerting-v3-policies-policy-id-14b9fe60
path: operations/notification-policies
description: Get details for a single policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies/{policy_id}
operation_ids:
    - notification-policies-get-a-notification-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Notification policy

`GET /accounts/{account_id}/alerting/v3/policies/{policy_id}`

Operation ID: `notification-policies-get-a-notification-policy`

Get details for a single policy.

## Definition

```yaml
{"operationId": "notification-policies-get-a-notification-policy", "summary": "Get a Notification policy", "description": "Get details for a single policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_policy-id"}}], "responses": {"200": {"description": "Get a Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_single_response"}}}}, "4XX": {"description": "Get a Notification policy response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_single_response"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification policies"], "x-api-token-group": ["Zero Trust: PII Read", "Notifications Write", "Notifications Read", "Account Settings Write", "Account Settings Read"]}
```
