---
title: Retrieve a transfer
page_id: operation-get-v1-transfers-transfer-7220d148
path: operations/untagged
description: <p>Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/transfers/{transfer}
operation_ids:
    - GetTransfersTransfer
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a transfer

`GET /v1/transfers/{transfer}`

Operation ID: `GetTransfersTransfer`

<p>Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.</p>

## Definition

```yaml
{"summary": "Retrieve a transfer", "description": "<p>Retrieves the details of an existing transfer. Supply the unique transfer ID from either a transfer creation request or the transfer list, and Stripe will return the corresponding transfer information.</p>", "operationId": "GetTransfersTransfer", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "transfer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/transfer"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
