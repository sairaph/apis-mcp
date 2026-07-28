---
title: Batch move accounts
page_id: operation-post-accounts-move-4d9a4b0d
path: operations/accounts
description: Batch move a collection of accounts to a specific organization. ⚠️ Not implemented.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/move
operation_ids:
    - Accounts_batchMoveAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch move accounts

`POST /accounts/move`

Operation ID: `Accounts_batchMoveAccounts`

Batch move a collection of accounts to a specific organization. ⚠️ Not implemented.

## Definition

```yaml
{"operationId": "Accounts_batchMoveAccounts", "summary": "Batch move accounts", "description": "Batch move a collection of accounts to a specific organization. ⚠️ Not implemented.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"account_ids": {"description": "Move these accounts to the destination organization.", "type": "array", "items": {"type": "string"}}, "destination_organization_id": {"description": "Move accounts to this organization ID.", "type": "string"}}, "required": ["account_ids", "destination_organization_id"]}}}}, "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_BatchAccountMoveResponse"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts", "x-fern-sdk-method-name": "move"}
```
