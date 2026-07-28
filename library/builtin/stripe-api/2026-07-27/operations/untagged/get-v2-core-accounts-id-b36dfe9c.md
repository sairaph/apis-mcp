---
title: Retrieve an account
page_id: operation-get-v2-core-accounts-id-af2e42b2
path: operations/untagged
description: Retrieves the details of an Account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/accounts/{id}
operation_ids:
    - GetV2CoreAccountsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve an account

`GET /v2/core/accounts/{id}`

Operation ID: `GetV2CoreAccountsId`

Retrieves the details of an Account.

## Definition

```yaml
{"summary": "Retrieve an account", "description": "Retrieves the details of an Account.", "operationId": "GetV2CoreAccountsId", "parameters": [{"name": "id", "in": "path", "description": "The ID of the Account to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "include", "in": "query", "description": "Additional fields to include in the response.", "required": false, "style": "deepObject", "schema": {"type": "array", "items": {"type": "string", "enum": ["configuration.customer", "configuration.merchant", "configuration.recipient", "defaults", "future_requirements", "identity", "requirements"]}}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.account"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_not_yet_compatible_with_v2"}, {"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error.v1_account_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error.v1_customer_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
