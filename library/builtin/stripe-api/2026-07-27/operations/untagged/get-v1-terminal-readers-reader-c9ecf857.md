---
title: Retrieve a Reader
page_id: operation-get-v1-terminal-readers-reader-f507027e
path: operations/untagged
description: <p>Retrieves a <code>Reader</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/terminal/readers/{reader}
operation_ids:
    - GetTerminalReadersReader
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Reader

`GET /v1/terminal/readers/{reader}`

Operation ID: `GetTerminalReadersReader`

<p>Retrieves a <code>Reader</code> object.</p>

## Definition

```yaml
{"summary": "Retrieve a Reader", "description": "<p>Retrieves a <code>Reader</code> object.</p>", "operationId": "GetTerminalReadersReader", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"anyOf": [{"$ref": "#/components/schemas/terminal.reader"}, {"$ref": "#/components/schemas/deleted_terminal.reader"}]}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
