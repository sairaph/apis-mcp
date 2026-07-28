---
title: Get a user
page_id: operation-get-accounts-account-id-access-users-user-id-1cf4fee9
path: operations/zero-trust-users
description: Gets a specific user for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}
operation_ids:
    - zero-trust-users-get-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a user

`GET /accounts/{account_id}/access/users/{user_id}`

Operation ID: `zero-trust-users-get-user`

Gets a specific user for an account.

## Definition

```yaml
{"operationId": "zero-trust-users-get-user", "summary": "Get a user", "description": "Gets a specific user for an account.", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_user_response"}}}}, "4XX": {"description": "Get user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": ["Access: Audit Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
