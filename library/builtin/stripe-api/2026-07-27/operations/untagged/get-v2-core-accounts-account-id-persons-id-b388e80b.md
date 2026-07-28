---
title: Retrieve a person
page_id: operation-get-v2-core-accounts-account-id-persons-id-55996c48
path: operations/untagged
description: Retrieves a Person associated with an Account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/accounts/{account_id}/persons/{id}
operation_ids:
    - GetV2CoreAccountsAccountIdPersonsId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a person

`GET /v2/core/accounts/{account_id}/persons/{id}`

Operation ID: `GetV2CoreAccountsAccountIdPersonsId`

Retrieves a Person associated with an Account.

## Definition

```yaml
{"summary": "Retrieve a person", "description": "Retrieves a Person associated with an Account.", "operationId": "GetV2CoreAccountsAccountIdPersonsId", "parameters": [{"name": "account_id", "in": "path", "description": "The Account the Person is associated with.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "id", "in": "path", "description": "The ID of the Person to retrieve.", "required": true, "style": "simple", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/v2.core.account_person"}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_not_yet_compatible_with_v2"}, {"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.not_found"}, {"$ref": "#/components/schemas/v2.error.v1_account_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error.v1_customer_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
