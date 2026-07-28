---
title: Retrieve a person token
page_id: operation-get-v2-core-accounts-account-id-person-tokens-id-b061ed13
path: operations/untagged
description: Retrieves a Person Token associated with an Account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/accounts/{account_id}/person_tokens/{id}
operation_ids:
    - GetV2CoreAccountsAccountIdPersonTokensId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a person token

`GET /v2/core/accounts/{account_id}/person_tokens/{id}`

Operation ID: `GetV2CoreAccountsAccountIdPersonTokensId`

Retrieves a Person Token associated with an Account.

## Definition

```yaml
{"summary": "Retrieve a person token", "description": "Retrieves a Person Token associated with an Account.", "operationId": "GetV2CoreAccountsAccountIdPersonTokensId", "parameters": [{"name": "account_id", "in": "path", "description": "The Account the Person is associated with.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "id", "in": "path", "description": "The ID of the Person Token to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.account_person_token"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
