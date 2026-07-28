---
title: List Tokens
page_id: operation-get-user-tokens-f8510772
path: operations/user-api-tokens
description: List all access tokens you created. Results include active, disabled, and recently-expired tokens when include_expired is set to true.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/tokens
operation_ids:
    - user-api-tokens-list-tokens
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Tokens

`GET /user/tokens`

Operation ID: `user-api-tokens-list-tokens`

List all access tokens you created. Results include active, disabled, and recently-expired tokens when include_expired is set to true.

## Definition

```yaml
{"operationId": "user-api-tokens-list-tokens", "summary": "List Tokens", "description": "List all access tokens you created. Results include active, disabled, and recently-expired tokens when include_expired is set to true.", "parameters": [{"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}, {"name": "direction", "in": "query", "schema": {"description": "Direction to order results.", "type": "string", "example": "desc", "enum": ["asc", "desc"]}}, {"$ref": "#/components/parameters/iam_include_expired"}], "responses": {"200": {"description": "List Tokens response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_tokens_response"}}}}, "4XX": {"description": "List Tokens response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-api-token-group": ["API Tokens Write", "API Tokens Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.token.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
