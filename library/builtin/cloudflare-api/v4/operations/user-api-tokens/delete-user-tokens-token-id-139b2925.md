---
title: Delete Token
page_id: operation-delete-user-tokens-token-id-ae928031
path: operations/user-api-tokens
description: Destroy a token.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /user/tokens/{token_id}
operation_ids:
    - user-api-tokens-delete-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Token

`DELETE /user/tokens/{token_id}`

Operation ID: `user-api-tokens-delete-token`

Destroy a token.

## Definition

```yaml
{"operationId": "user-api-tokens-delete-token", "summary": "Delete Token", "description": "Destroy a token.", "parameters": [{"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-single-id"}}}}, "4XX": {"description": "Delete Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.delete"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
