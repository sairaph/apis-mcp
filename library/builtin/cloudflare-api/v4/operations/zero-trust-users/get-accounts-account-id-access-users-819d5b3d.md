---
title: Get users
page_id: operation-get-accounts-account-id-access-users-8e83c696
path: operations/zero-trust-users
description: Gets a list of users for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/users
operation_ids:
    - zero-trust-users-get-users
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get users

`GET /accounts/{account_id}/access/users`

Operation ID: `zero-trust-users-get-users`

Gets a list of users for an account.

## Definition

```yaml
{"operationId": "zero-trust-users-get-users", "summary": "Get users", "description": "Gets a list of users for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "The name of the user.", "type": "string"}}, {"name": "email", "in": "query", "schema": {"description": "The email of the user.", "type": "string"}}, {"name": "search", "in": "query", "schema": {"description": "Search for users by other listed query parameters.", "type": "string"}}, {"$ref": "#/components/parameters/access_page"}, {"$ref": "#/components/parameters/access_per_page"}], "responses": {"200": {"description": "Get users response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-24"}}}}, "4XX": {"description": "Get users response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": ["Access: Audit Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
