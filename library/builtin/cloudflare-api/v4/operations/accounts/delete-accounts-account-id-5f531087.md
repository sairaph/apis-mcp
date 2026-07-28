---
title: Delete a specific account
page_id: operation-delete-accounts-account-id-8cac32ca
path: operations/accounts
description: Delete a specific account (only available for tenant admins at this time). This is a permanent operation that will delete any zones or other resources under the account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}
operation_ids:
    - account-deletion
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a specific account

`DELETE /accounts/{account_id}`

Operation ID: `account-deletion`

Delete a specific account (only available for tenant admins at this time). This is a permanent operation that will delete any zones or other resources under the account

## Definition

```yaml
{"operationId": "account-deletion", "summary": "Delete a specific account", "description": "Delete a specific account (only available for tenant admins at this time). This is a permanent operation that will delete any zones or other resources under the account", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"description": "The account ID of the account to be deleted", "type": "string"}}], "responses": {"200": {"description": "Account Deletion Success Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Account Deletion Failure Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": ["Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": false, "enterprise": false, "free": false, "pro": false}}
```
