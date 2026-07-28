---
title: Update a user
page_id: operation-put-accounts-account-id-access-users-user-id-1a158894
path: operations/zero-trust-users
description: Updates a specific user's name for an account. Requires the user's current email as confirmation (email cannot be changed).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}
operation_ids:
    - zero-trust-users-update-user
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a user

`PUT /accounts/{account_id}/access/users/{user_id}`

Operation ID: `zero-trust-users-update-user`

Updates a specific user's name for an account. Requires the user's current email as confirmation (email cannot be changed).

## Definition

```yaml
{"operationId": "zero-trust-users-update-user", "summary": "Update a user", "description": "Updates a specific user's name for an account. Requires the user's current email as confirmation (email cannot be changed).", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"email": {"$ref": "#/components/schemas/access_email-2"}, "name": {"$ref": "#/components/schemas/access_name-10"}}, "required": ["name", "email"]}}}}, "responses": {"200": {"description": "Update user response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_user_response"}}}}, "4XX": {"description": "Update user response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": ["Zero Trust: Seats Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
