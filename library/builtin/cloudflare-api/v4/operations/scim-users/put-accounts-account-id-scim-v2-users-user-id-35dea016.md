---
title: Replace SCIM User
page_id: operation-put-accounts-account-id-scim-v2-users-user-id-9ead0e49
path: operations/scim-users
description: Replaces a SCIM User resource (RFC 7644 Section 3.5.1). Fully replaces the mutable attributes of the user. Supports updating `userName`, `name`, `emails`, and `active`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/scim/v2/Users/{user_id}
operation_ids:
    - scim-users-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace SCIM User

`PUT /accounts/{account_id}/scim/v2/Users/{user_id}`

Operation ID: `scim-users-put`

Replaces a SCIM User resource (RFC 7644 Section 3.5.1). Fully replaces the mutable attributes of the user. Supports updating `userName`, `name`, `emails`, and `active`.

## Definition

```yaml
{"operationId": "scim-users-put", "summary": "Replace SCIM User", "description": "Replaces a SCIM User resource (RFC 7644 Section 3.5.1). Fully replaces the mutable attributes of the user. Supports updating `userName`, `name`, `emails`, and `active`.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_scim_user_identifier"}}], "requestBody": {"required": true, "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user_replace_request"}}}}, "responses": {"200": {"description": "Replace SCIM User response", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_user"}}}}, "4XX": {"description": "Replace SCIM User response failure", "content": {"application/scim+json": {"schema": {"$ref": "#/components/schemas/iam_scim_error_response"}}}}}, "security": [{"api_token": []}], "tags": ["SCIM Users"], "x-api-token-group": ["SCIM Provisioning", "Account Settings Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.account.member.update"]}, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
