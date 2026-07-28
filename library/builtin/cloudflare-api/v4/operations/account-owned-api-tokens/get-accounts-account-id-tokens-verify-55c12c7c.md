---
title: Verify Token
page_id: operation-get-accounts-account-id-tokens-verify-3a284014
path: operations/account-owned-api-tokens
description: Test whether a token works.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tokens/verify
operation_ids:
    - account-api-tokens-verify-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Verify Token

`GET /accounts/{account_id}/tokens/verify`

Operation ID: `account-api-tokens-verify-token`

Test whether a token works.

## Definition

```yaml
{"operationId": "account-api-tokens-verify-token", "summary": "Verify Token", "description": "Test whether a token works.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}], "responses": {"200": {"description": "Verify Token response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_token_verify_response_single_segment"}}}}, "4XX": {"description": "Verify Token response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}], "tags": ["Account Owned API Tokens"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
