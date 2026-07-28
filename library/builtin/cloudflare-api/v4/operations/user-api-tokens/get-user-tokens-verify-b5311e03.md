---
title: Verify Token
page_id: operation-get-user-tokens-verify-554edb05
path: operations/user-api-tokens
description: Test whether a token works.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /user/tokens/verify
operation_ids:
    - user-api-tokens-verify-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Verify Token

`GET /user/tokens/verify`

Operation ID: `user-api-tokens-verify-token`

Test whether a token works.

## Definition

```yaml
{"operationId": "user-api-tokens-verify-token", "summary": "Verify Token", "description": "Test whether a token works.", "responses": {"200": {"description": "Verify Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_token_verify_response_single_segment"}}}}, "4XX": {"description": "Verify Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["User API Tokens"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
