---
title: Get SCIM User
page_id: operation-get-accounts-account-id-scim-v2-users-user-id-61df6745
path: operations/scim-users
description: Retrieves a single account member as a SCIM User resource by user tag.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/scim/v2/Users/{user_id}
operation_ids:
    - scim-users-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get SCIM User

`GET /accounts/{account_id}/scim/v2/Users/{user_id}`

Operation ID: `scim-users-get`

Retrieves a single account member as a SCIM User resource by user tag.

## Definition

```yaml
{"operationId": "scim-users-get", "summary": "Get SCIM User", "description": "Retrieves a single account member as a SCIM User resource by user tag.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_user_identifier"}}], "responses": {"200": {"description": "Get SCIM User response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user"}}}}, "4XX": {"description": "Get SCIM User response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Users"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
