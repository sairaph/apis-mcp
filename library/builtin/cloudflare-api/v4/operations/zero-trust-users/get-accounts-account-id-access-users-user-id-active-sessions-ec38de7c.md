---
title: Get active sessions
page_id: operation-get-accounts-account-id-access-users-user-id-active-sessions-c5c58324
path: operations/zero-trust-users
description: Get active sessions for a single user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/users/{user_id}/active_sessions
operation_ids:
    - zero-trust-users-get-active-sessions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get active sessions

`GET /accounts/{account_id}/access/users/{user_id}/active_sessions`

Operation ID: `zero-trust-users-get-active-sessions`

Get active sessions for a single user.

## Definition

```yaml
{"operationId": "zero-trust-users-get-active-sessions", "summary": "Get active sessions", "description": "Get active sessions for a single user.", "parameters": [{"name": "user_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get active sessions response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_active_sessions_response"}}}}, "4XX": {"description": "Get active sessions response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust users"], "x-api-token-group": ["Access: Audit Logs Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.users.active-sessions", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
