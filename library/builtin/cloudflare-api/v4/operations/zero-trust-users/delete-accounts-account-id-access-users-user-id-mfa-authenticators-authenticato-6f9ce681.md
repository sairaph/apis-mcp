---
title: Delete a user's MFA device
page_id: operation-delete-accounts-account-id-access-users-user-id-mfa-authenticators-authe-f4bad20c
path: operations/zero-trust-users
description: Deletes a specific MFA device for a user. This action is only available if MFA is turned on for the organization.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}/mfa_authenticators/{authenticator_id}
operation_ids:
    - zero-trust-users-delete-mfa-authenticator
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a user's MFA device

`DELETE /accounts/{account_id}/access/users/{user_id}/mfa_authenticators/{authenticator_id}`

Operation ID: `zero-trust-users-delete-mfa-authenticator`

Deletes a specific MFA device for a user. This action is only available if MFA is turned on for the organization.

## Definition

```yaml
{"operationId": "zero-trust-users-delete-mfa-authenticator", "summary": "Delete a user's MFA device", "description": "Deletes a specific MFA device for a user. This action is only available if MFA is turned on for the organization.", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "authenticator_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_authenticator_id"}}], "responses": {"200": {"description": "Delete authenticator response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_delete_authenticator_response"}}}}, "4XX": {"description": "Delete authenticator response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.users.mfa.authenticators", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
