---
title: Set reader display
page_id: operation-post-v1-terminal-readers-reader-set-reader-display-ee3b2a71
path: operations/untagged
description: <p>Sets the reader display to show <a href="/docs/terminal/features/display">cart details</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/set_reader_display
operation_ids:
    - PostTerminalReadersReaderSetReaderDisplay
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Set reader display

`POST /v1/terminal/readers/{reader}/set_reader_display`

Operation ID: `PostTerminalReadersReaderSetReaderDisplay`

<p>Sets the reader display to show <a href="/docs/terminal/features/display">cart details</a>.</p>

## Definition

```yaml
{"summary": "Set reader display", "description": "<p>Sets the reader display to show <a href=\"/docs/terminal/features/display\">cart details</a>.</p>", "operationId": "PostTerminalReadersReaderSetReaderDisplay", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["type"], "type": "object", "properties": {"cart": {"title": "cart", "required": ["currency", "line_items", "total"], "type": "object", "properties": {"currency": {"type": "string", "format": "currency"}, "line_items": {"type": "array", "items": {"title": "line_item", "required": ["amount", "description", "quantity"], "type": "object", "properties": {"amount": {"type": "integer"}, "description": {"maxLength": 5000, "type": "string"}, "quantity": {"type": "integer"}}}}, "tax": {"type": "integer"}, "total": {"type": "integer"}}, "description": "Cart details to display on the reader screen, including line items, amounts, and currency."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "type": {"type": "string", "description": "Type of information to display. Only `cart` is currently supported.", "enum": ["cart"]}}, "additionalProperties": false}, "encoding": {"cart": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
