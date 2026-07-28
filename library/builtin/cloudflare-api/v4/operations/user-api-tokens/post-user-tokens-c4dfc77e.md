---
title: Create Token
page_id: operation-post-user-tokens-16b0eecd
path: operations/user-api-tokens
description: Create a new access token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /user/tokens
operation_ids:
    - user-api-tokens-create-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Token

`POST /user/tokens`

Operation ID: `user-api-tokens-create-token`

Create a new access token.

## Definition

```yaml
{"operationId": "user-api-tokens-create-token", "summary": "Create Token", "description": "Create a new access token.", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_create_payload"}}}}, "responses": {"200": {"description": "Create Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_create_response"}}}}, "4XX": {"description": "Create Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.create"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
