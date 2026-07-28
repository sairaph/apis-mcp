---
title: Retrieve a TransactionEntry
page_id: operation-get-v1-treasury-transaction-entries-id-11bd4f27
path: operations/untagged
description: <p>Retrieves a TransactionEntry object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/treasury/transaction_entries/{id}
operation_ids:
    - GetTreasuryTransactionEntriesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a TransactionEntry

`GET /v1/treasury/transaction_entries/{id}`

Operation ID: `GetTreasuryTransactionEntriesId`

<p>Retrieves a TransactionEntry object.</p>

## Definition

```yaml
{"summary": "Retrieve a TransactionEntry", "description": "<p>Retrieves a TransactionEntry object.</p>", "operationId": "GetTreasuryTransactionEntriesId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.transaction_entry"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
