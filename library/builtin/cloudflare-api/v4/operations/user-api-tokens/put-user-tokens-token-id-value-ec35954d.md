---
title: Roll Token
page_id: operation-put-user-tokens-token-id-value-b3199295
path: operations/user-api-tokens
description: Roll the token secret.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /user/tokens/{token_id}/value
operation_ids:
    - user-api-tokens-roll-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Roll Token

`PUT /user/tokens/{token_id}/value`

Operation ID: `user-api-tokens-roll-token`

Roll the token secret.

## Definition

```yaml
{"operationId": "user-api-tokens-roll-token", "summary": "Roll Token", "description": "Roll the token secret.", "parameters": [{"name": "token_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_token_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object"}}}}, "responses": {"200": {"description": "Roll Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_response_single_value"}}}}, "4XX": {"description": "Roll Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.update"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
