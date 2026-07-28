---
title: Token Details
page_id: operation-get-user-tokens-token-id-64e95553
path: operations/user-api-tokens
description: Get information about a specific token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/tokens/{token_id}
operation_ids:
    - user-api-tokens-token-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Token Details

`GET /user/tokens/{token_id}`

Operation ID: `user-api-tokens-token-details`

Get information about a specific token.

## Definition

```yaml
{"operationId": "user-api-tokens-token-details", "summary": "Token Details", "description": "Get information about a specific token.", "parameters": [{"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "responses": {"200": {"description": "Token Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_response"}}}}, "4XX": {"description": "Token Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write", "API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
