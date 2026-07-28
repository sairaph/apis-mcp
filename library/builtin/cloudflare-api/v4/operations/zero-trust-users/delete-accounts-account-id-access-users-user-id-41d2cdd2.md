---
title: Delete a user
page_id: operation-delete-accounts-account-id-access-users-user-id-5dac73f7
path: operations/zero-trust-users
description: Deletes a specific user for an account. This will also revoke any active seats and tokens for the user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}
operation_ids:
    - zero-trust-users-delete-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a user

`DELETE /accounts/{account_id}/access/users/{user_id}`

Operation ID: `zero-trust-users-delete-user`

Deletes a specific user for an account. This will also revoke any active seats and tokens for the user.

## Definition

```yaml
{"operationId": "zero-trust-users-delete-user", "summary": "Delete a user", "description": "Deletes a specific user for an account. This will also revoke any active seats and tokens for the user.", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_delete_user_response"}}}}, "4XX": {"description": "Delete user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
