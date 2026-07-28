---
title: Simulate a successful input collection
page_id: operation-post-v1-test-helpers-terminal-readers-reader-succeed-input-collection-6a98279e
path: operations/untagged
description: <p>Use this endpoint to trigger a successful input collection on a simulated reader.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/terminal/readers/{reader}/succeed_input_collection
operation_ids:
    - PostTestHelpersTerminalReadersReaderSucceedInputCollection
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Simulate a successful input collection

`POST /v1/test_helpers/terminal/readers/{reader}/succeed_input_collection`

Operation ID: `PostTestHelpersTerminalReadersReaderSucceedInputCollection`

<p>Use this endpoint to trigger a successful input collection on a simulated reader.</p>

## Definition

```yaml
{"summary": "Simulate a successful input collection", "description": "<p>Use this endpoint to trigger a successful input collection on a simulated reader.</p>", "operationId": "PostTestHelpersTerminalReadersReaderSucceedInputCollection", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "skip_non_required_inputs": {"type": "string", "description": "This parameter defines the skip behavior for input collection.", "enum": ["all", "none"]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
