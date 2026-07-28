---
title: List all FinancialAccounts
page_id: operation-get-v1-treasury-financial-accounts-a1b5c696
path: operations/untagged
description: <p>Returns a list of FinancialAccounts.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/treasury/financial_accounts
operation_ids:
    - GetTreasuryFinancialAccounts
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List all FinancialAccounts

`GET /v1/treasury/financial_accounts`

Operation ID: `GetTreasuryFinancialAccounts`

<p>Returns a list of FinancialAccounts.</p>

## Definition

```yaml
{"summary": "List all FinancialAccounts", "description": "<p>Returns a list of FinancialAccounts.</p>", "operationId": "GetTreasuryFinancialAccounts", "parameters": [{"name": "created", "in": "query", "description": "Only return FinancialAccounts that were created during the given date interval.", "required": false, "style": "deepObject", "explode": true, "schema": {"anyOf": [{"title": "range_query_specs", "type": "object", "properties": {"gt": {"type": "integer"}, "gte": {"type": "integer"}, "lt": {"type": "integer"}, "lte": {"type": "integer"}}}, {"type": "integer"}]}}, {"name": "ending_before", "in": "query", "description": "An object ID cursor for use in pagination.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit ranging from 1 to 100 (defaults to 10).", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "starting_after", "in": "query", "description": "An object ID cursor for use in pagination.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "status", "in": "query", "description": "Only return FinancialAccounts that have the given status: `open` or `closed`", "required": false, "style": "form", "explode": true, "schema": {"type": "string", "enum": ["closed", "open"], "x-stripeBypassValidation": true}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "TreasuryFinancialAccountsResourceFinancialAccountList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "items": {"$ref": "#/components/schemas/treasury.financial_account"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "pattern": "^/v1/treasury/financial_accounts", "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
