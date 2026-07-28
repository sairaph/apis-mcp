---
title: Retrieve balance
page_id: operation-get-v1-balance-e8c72cba
path: operations/untagged
description: |-
    <p>Retrieves the current account balance, based on the authentication that was used to make the request.
     For a sample request, see <a href="/docs/connect/account-balances#accounting-for-negative-balances">Accounting for negative balances</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/balance
operation_ids:
    - GetBalance
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve balance

`GET /v1/balance`

Operation ID: `GetBalance`

<p>Retrieves the current account balance, based on the authentication that was used to make the request.
 For a sample request, see <a href="/docs/connect/account-balances#accounting-for-negative-balances">Accounting for negative balances</a>.</p>

## Definition

```yaml
{"summary": "Retrieve balance", "description": "<p>Retrieves the current account balance, based on the authentication that was used to make the request.\n For a sample request, see <a href=\"/docs/connect/account-balances#accounting-for-negative-balances\">Accounting for negative balances</a>.</p>", "operationId": "GetBalance", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/balance"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
