---
title: Create an account
page_id: operation-post-accounts-6666e693
path: operations/accounts
description: Create an account (only available for tenant admins at this time)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts
operation_ids:
    - account-creation
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an account

`POST /accounts`

Operation ID: `account-creation`

Create an account (only available for tenant admins at this time)

## Definition

```yaml
{"operationId": "account-creation", "summary": "Create an account", "description": "Create an account (only available for tenant admins at this time)", "requestBody": {"description": "Parameters for account creation", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_create-account"}}}}, "responses": {"200": {"description": "Account Creation Success Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_response_single_account"}}}}, "4XX": {"description": "Account Creation Failure Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-cfPlanAvailability": {"business": false, "enterprise": false, "free": false, "pro": false}}
```
