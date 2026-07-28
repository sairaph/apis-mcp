---
title: List persons
page_id: operation-get-v2-core-accounts-account-id-persons-0214fc12
path: operations/untagged
description: Returns a paginated list of Persons associated with an Account.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/accounts/{account_id}/persons
operation_ids:
    - GetV2CoreAccountsAccountIdPersons
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List persons

`GET /v2/core/accounts/{account_id}/persons`

Operation ID: `GetV2CoreAccountsAccountIdPersons`

Returns a paginated list of Persons associated with an Account.

## Definition

```yaml
{"summary": "List persons", "description": "Returns a paginated list of Persons associated with an Account.", "operationId": "GetV2CoreAccountsAccountIdPersons", "parameters": [{"name": "account_id", "in": "path", "description": "Account the Persons are associated with.", "required": true, "style": "simple", "schema": {"type": "string"}}, {"name": "limit", "in": "query", "description": "The upper limit on the number of accounts returned by the List Account request.", "required": false, "style": "form", "schema": {"type": "integer"}}, {"name": "page", "in": "query", "description": "The page token to navigate to next or previous batch of accounts given by the list request.", "required": false, "style": "form", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"required": ["data", "next_page_url", "previous_page_url"], "type": "object", "properties": {"data": {"type": "array", "description": "A list of retrieved Person objects.", "items": {"$ref": "#/components/schemas/v2.core.account_person"}}, "next_page_url": {"type": "string", "description": "URL to fetch the next page of the list. If there are no more pages, the value is null.", "nullable": true}, "previous_page_url": {"type": "string", "description": "URL to fetch the previous page of the list. If there are no previous pages, the value is null.", "nullable": true}}}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_not_yet_compatible_with_v2"}, {"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.v1_account_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error.v1_customer_instead_of_v2_account"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
