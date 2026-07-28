---
title: Get last seen identity
page_id: operation-get-accounts-account-id-access-users-user-id-last-seen-identity-ab5a9ad9
path: operations/zero-trust-users
description: Get last seen identity for a single user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}/last_seen_identity
operation_ids:
    - zero-trust-users-get-last-seen-identity
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get last seen identity

`GET /accounts/{account_id}/access/users/{user_id}/last_seen_identity`

Operation ID: `zero-trust-users-get-last-seen-identity`

Get last seen identity for a single user.

## Definition

```yaml
{"operationId": "zero-trust-users-get-last-seen-identity", "summary": "Get last seen identity", "description": "Get last seen identity for a single user.", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get active session response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_last_seen_identity_response"}}}}, "4XX": {"description": "Get active session response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": ["Access: Audit Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users.last-seen-identity", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
