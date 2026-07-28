---
title: Update Token
page_id: operation-put-user-tokens-token-id-743b4d86
path: operations/user-api-tokens
description: Update an existing token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /user/tokens/{token_id}
operation_ids:
    - user-api-tokens-update-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Token

`PUT /user/tokens/{token_id}`

Operation ID: `user-api-tokens-update-token`

Update an existing token.

## Definition

```yaml
{"operationId": "user-api-tokens-update-token", "summary": "Update Token", "description": "Update an existing token.", "parameters": [{"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_token_body"}}}}, "responses": {"200": {"description": "Update Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_token_response"}}}}, "4XX": {"description": "Update Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
