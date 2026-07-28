---
title: Move account
page_id: operation-post-accounts-account-id-move-f6184bfa
path: operations/accounts
description: Move an account within an organization hierarchy or an account outside an organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/move
operation_ids:
    - Accounts_moveAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Move account

`POST /accounts/{account_id}/move`

Operation ID: `Accounts_moveAccounts`

Move an account within an organization hierarchy or an account outside an organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)

## Definition

```yaml
{"operationId": "Accounts_moveAccounts", "summary": "Move account", "description": "Move an account within an organization hierarchy or an account outside an organization. (Currently in Public Beta - see https://developers.cloudflare.com/fundamentals/organizations/)", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "The destination organization ID is where the account is to be moved.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"destination_organization_id": {"type": "string"}}, "required": ["destination_organization_id"]}}}}, "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_MoveAccountResponse"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.account-organizations", "x-fern-sdk-method-name": "create"}
```
