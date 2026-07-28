---
title: Delete a Notification policy
page_id: operation-delete-accounts-account-id-alerting-v3-policies-policy-id-a591ab6d
path: operations/notification-policies
description: Delete a Notification policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/alerting/v3/policies/{policy_id}
operation_ids:
    - notification-policies-delete-a-notification-policy
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Notification policy

`DELETE /accounts/{account_id}/alerting/v3/policies/{policy_id}`

Operation ID: `notification-policies-delete-a-notification-policy`

Delete a Notification policy.

## Definition

```yaml
{"operationId": "notification-policies-delete-a-notification-policy", "summary": "Delete a Notification policy", "description": "Delete a Notification policy.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_account-id"}}, {"name": "policy_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/aaa_policy-id"}}], "responses": {"200": {"description": "Delete a Notification policy response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/aaa_api-response-collection"}}}}, "4XX": {"description": "Delete a Notification policy response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/aaa_api-response-collection"}, {"$ref": "#/components/schemas/aaa_api-response-common-failure-3"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Notification policies"], "x-api-token-group": ["Notifications Write", "Account Settings Write"]}
```
