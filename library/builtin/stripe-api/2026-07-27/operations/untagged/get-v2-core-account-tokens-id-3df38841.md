---
title: Retrieve an account token
page_id: operation-get-v2-core-account-tokens-id-43b09340
path: operations/untagged
description: Retrieves an Account Token.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/account_tokens/{id}
operation_ids:
    - GetV2CoreAccountTokensId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an account token

`GET /v2/core/account_tokens/{id}`

Operation ID: `GetV2CoreAccountTokensId`

Retrieves an Account Token.

## Definition

```yaml
{"summary": "Retrieve an account token", "description": "Retrieves an Account Token.", "operationId": "GetV2CoreAccountTokensId", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Account Token to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.account_token"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
