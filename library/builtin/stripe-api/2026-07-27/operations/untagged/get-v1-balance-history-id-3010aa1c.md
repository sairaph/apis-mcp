---
title: Retrieve a balance transaction
page_id: operation-get-v1-balance-history-id-ae11ed21
path: operations/untagged
description: |-
    <p>Retrieves the balance transaction with the given ID.</p>

    <p>Note that this endpoint previously used the path <code>/v1/balance/history/:id</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/balance/history/{id}
operation_ids:
    - GetBalanceHistoryId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a balance transaction

`GET /v1/balance/history/{id}`

Operation ID: `GetBalanceHistoryId`

<p>Retrieves the balance transaction with the given ID.</p>

<p>Note that this endpoint previously used the path <code>/v1/balance/history/:id</code>.</p>

## Definition

```yaml
{"summary": "Retrieve a balance transaction", "description": "<p>Retrieves the balance transaction with the given ID.</p>\n\n<p>Note that this endpoint previously used the path <code>/v1/balance/history/:id</code>.</p>", "operationId": "GetBalanceHistoryId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/balance_transaction"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
