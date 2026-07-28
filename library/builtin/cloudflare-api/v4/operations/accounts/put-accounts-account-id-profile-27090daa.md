---
title: Modify account profile
page_id: operation-put-accounts-account-id-profile-d72350aa
path: operations/accounts
description: Updates the profile information for a Cloudflare account. Allows modification of account-level settings and organizational details. Requires Account Settings Write permission.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/profile
operation_ids:
    - Accounts_modifyAccountProfile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Modify account profile

`PUT /accounts/{account_id}/profile`

Operation ID: `Accounts_modifyAccountProfile`

Updates the profile information for a Cloudflare account. Allows modification of account-level settings and organizational details. Requires Account Settings Write permission.

## Definition

```yaml
{"operationId": "Accounts_modifyAccountProfile", "summary": "Modify account profile", "description": "Updates the profile information for a Cloudflare account. Allows modification of account-level settings and organizational details. Requires Account Settings Write permission.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_Profile"}}}}, "responses": {"204": {"description": "There is no content to send for this request, but the headers may be useful."}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": ["Account Settings Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.account-profile", "x-fern-sdk-method-name": "update"}
```
