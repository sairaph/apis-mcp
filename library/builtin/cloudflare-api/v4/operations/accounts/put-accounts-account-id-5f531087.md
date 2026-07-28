---
title: Update Account
page_id: operation-put-accounts-account-id-86909bfc
path: operations/accounts
description: Update an existing account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}
operation_ids:
    - accounts-update-account
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Account

`PUT /accounts/{account_id}`

Operation ID: `accounts-update-account`

Update an existing account.

## Definition

```yaml
{"operationId": "accounts-update-account", "summary": "Update Account", "description": "Update an existing account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_components-schemas-account"}}}}, "responses": {"200": {"description": "Update Account response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_response_single_account"}}}}, "4XX": {"description": "Update Account response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": ["Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["#organization:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
