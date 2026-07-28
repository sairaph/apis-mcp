---
title: Create a user
page_id: operation-post-accounts-account-id-access-users-9edc29c4
path: operations/zero-trust-users
description: Creates a new user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/users
operation_ids:
    - zero-trust-users-create-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a user

`POST /accounts/{account_id}/access/users`

Operation ID: `zero-trust-users-create-user`

Creates a new user.

## Definition

```yaml
{"operationId": "zero-trust-users-create-user", "summary": "Create a user", "description": "Creates a new user.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"email": {"$ref": "#/components/schemas/access_email-2"}, "name": {"$ref": "#/components/schemas/access_name-10"}}, "required": ["email"]}}}}, "responses": {"201": {"description": "Create user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_user_response"}}}}, "4XX": {"description": "Create user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
