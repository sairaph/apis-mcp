---
title: Revoke all Access tokens for a user
page_id: operation-post-accounts-account-id-access-organizations-revoke-user-f71d1028
path: operations/zero-trust-organization
description: Revokes a user's access across all applications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/organizations/revoke_user
operation_ids:
    - zero-trust-organization-revoke-all-access-tokens-for-a-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Revoke all Access tokens for a user

`POST /accounts/{account_id}/access/organizations/revoke_user`

Operation ID: `zero-trust-organization-revoke-all-access-tokens-for-a-user`

Revokes a user's access across all applications.

## Definition

```yaml
{"operationId": "zero-trust-organization-revoke-all-access-tokens-for-a-user", "summary": "Revoke all Access tokens for a user", "description": "Revokes a user's access across all applications.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "devices", "in": "query", "description": "When set to `true`, all devices associated with the user will be revoked.", "schema": {"type": "boolean"}, "example": true}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"devices": {"description": "When set to `true`, all devices associated with the user will be revoked.", "type": "boolean", "example": true}, "email": {"description": "The email of the user to revoke.", "type": "string", "example": "test@example.com"}, "user_uid": {"description": "The uuid of the user to revoke.", "type": "string", "example": "699d98642c564d2e855e9661899b7252"}, "warp_session_reauth": {"description": "When set to `true`, the user will be required to re-authenticate to WARP for all Gateway policies that enforce a WARP client session duration. When `false`, the user’s WARP session will remain active", "type": "boolean", "example": true}}, "required": ["email"]}}}}, "responses": {"200": {"description": "Revoke all Access tokens for a user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_empty_response"}}}}, "4xx": {"description": "Revoke all Access tokens for a user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust organization"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.organizations", "x-fern-sdk-method-name": "revoke-users", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation revokes all access tokens for a user destructively."}
```
