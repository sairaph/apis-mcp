---
title: Retrieve a DebitReversal
page_id: operation-get-v1-treasury-debit-reversals-debit-reversal-ab0422a5
path: operations/untagged
description: <p>Retrieves a DebitReversal object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/treasury/debit_reversals/{debit_reversal}
operation_ids:
    - GetTreasuryDebitReversalsDebitReversal
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a DebitReversal

`GET /v1/treasury/debit_reversals/{debit_reversal}`

Operation ID: `GetTreasuryDebitReversalsDebitReversal`

<p>Retrieves a DebitReversal object.</p>

## Definition

```yaml
{"summary": "Retrieve a DebitReversal", "description": "<p>Retrieves a DebitReversal object.</p>", "operationId": "GetTreasuryDebitReversalsDebitReversal", "parameters": [{"name": "debit_reversal", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/treasury.debit_reversal"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
