---
title: List Account Owners
page_id: operation-get-v1-financial-connections-accounts-account-owners-dd6b3910
path: operations/untagged
description: <p>Lists all owners for a given <code>Account</code></p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/financial_connections/accounts/{account}/owners
operation_ids:
    - GetFinancialConnectionsAccountsAccountOwners
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# List Account Owners

`GET /v1/financial_connections/accounts/{account}/owners`

Operation ID: `GetFinancialConnectionsAccountsAccountOwners`

<p>Lists all owners for a given <code>Account</code></p>

## Definition

```yaml
{"summary": "List Account Owners", "description": "<p>Lists all owners for a given <code>Account</code></p>", "operationId": "GetFinancialConnectionsAccountsAccountOwners", "parameters": [{"name": "account", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "ending_before", "in": "query", "description": "A cursor for use in pagination. `ending_before` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, starting with `obj_bar`, your subsequent call can include `ending_before=obj_bar` in order to fetch the previous page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "limit", "in": "query", "description": "A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.", "required": false, "style": "form", "explode": true, "schema": {"type": "integer"}}, {"name": "ownership", "in": "query", "description": "The ID of the ownership object to fetch owners from.", "required": true, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "starting_after", "in": "query", "description": "A cursor for use in pagination. `starting_after` is an object ID that defines your place in the list. For instance, if you make a list request and receive 100 objects, ending with `obj_foo`, your subsequent call can include `starting_after=obj_foo` in order to fetch the next page of the list.", "required": false, "style": "form", "explode": true, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"title": "BankConnectionsResourceOwnerList", "required": ["data", "has_more", "object", "url"], "type": "object", "properties": {"data": {"type": "array", "description": "Details about each object.", "items": {"$ref": "#/components/schemas/financial_connections.account_owner"}}, "has_more": {"type": "boolean", "description": "True if this list has another page of items after this one that can be fetched."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value. Always has the value `list`.", "enum": ["list"]}, "url": {"maxLength": 5000, "type": "string", "description": "The URL where this list can be accessed."}}, "description": "", "x-expandableFields": ["data"]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
