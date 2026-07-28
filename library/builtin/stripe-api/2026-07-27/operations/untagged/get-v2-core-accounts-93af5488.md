---
title: List accounts
page_id: operation-get-v2-core-accounts-a2858396
path: operations/untagged
description: Returns a list of Accounts.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v2/core/accounts
operation_ids:
    - GetV2CoreAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List accounts

`GET /v2/core/accounts`

Operation ID: `GetV2CoreAccounts`

Returns a list of Accounts.

## Definition

```yaml
{"summary": "List accounts", "description": "Returns a list of Accounts.", "operationId": "GetV2CoreAccounts", "parameters": [{"name": "applied_configurations", "in": "query", "description": "Filter only accounts that have all of the configurations specified. If omitted, returns all accounts regardless of which configurations they have.", "required": false, "style": "deepObject", "schema": {"type": "array", "items": {"type": "string", "enum": ["customer", "merchant", "recipient"]}}}, {"name": "closed", "in": "query", "description": "Filter by whether the account is closed. If omitted, returns only Accounts that are not closed.", "required": false, "style": "form", "schema": {"type": "boolean"}}, {"name": "limit", "in": "query", "description": "The upper limit on the number of accounts returned by the List Account request.", "required": false, "style": "form", "schema": {"type": "integer"}}, {"name": "page", "in": "query", "description": "The page token to navigate to next or previous batch of accounts given by the list request.", "required": false, "style": "form", "schema": {"type": "string"}}], "responses": {"200": {"description": "Successful response", "content": {"application/json": {"schema": {"required": ["data", "next_page_url", "previous_page_url"], "type": "object", "properties": {"data": {"type": "array", "description": "A list of retrieved Account objects.", "items": {"$ref": "#/components/schemas/v2.core.account"}}, "next_page_url": {"type": "string", "description": "URL to fetch the next page of the list. If there are no more pages, the value is null.", "nullable": true}, "previous_page_url": {"type": "string", "description": "URL to fetch the previous page of the list. If there are no previous pages, the value is null.", "nullable": true}}}}}}, "default": {"description": "Error response", "content": {"application/json": {"schema": {"oneOf": [{"$ref": "#/components/schemas/v2.error.account_rate_limit_exceeded"}, {"$ref": "#/components/schemas/v2.error.accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error.non_connect_platform_accounts_v2_access_blocked"}, {"$ref": "#/components/schemas/v2.error"}]}}}}}}
```
