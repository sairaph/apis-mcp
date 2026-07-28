---
title: Simulate an input collection timeout
page_id: operation-post-v1-test-helpers-terminal-readers-reader-timeout-input-collection-1779027d
path: operations/untagged
description: <p>Use this endpoint to complete an input collection with a timeout error on a simulated reader.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/terminal/readers/{reader}/timeout_input_collection
operation_ids:
    - PostTestHelpersTerminalReadersReaderTimeoutInputCollection
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Simulate an input collection timeout

`POST /v1/test_helpers/terminal/readers/{reader}/timeout_input_collection`

Operation ID: `PostTestHelpersTerminalReadersReaderTimeoutInputCollection`

<p>Use this endpoint to complete an input collection with a timeout error on a simulated reader.</p>

## Definition

```yaml
{"summary": "Simulate an input collection timeout", "description": "<p>Use this endpoint to complete an input collection with a timeout error on a simulated reader.</p>", "operationId": "PostTestHelpersTerminalReadersReaderTimeoutInputCollection", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
