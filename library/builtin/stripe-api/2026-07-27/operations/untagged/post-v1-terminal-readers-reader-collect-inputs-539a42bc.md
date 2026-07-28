---
title: Collect inputs using a Reader
page_id: operation-post-v1-terminal-readers-reader-collect-inputs-006803ae
path: operations/untagged
description: <p>Initiates an <a href="/docs/terminal/features/collect-inputs">input collection flow</a> on a Reader to display input forms and collect information from your customers.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/readers/{reader}/collect_inputs
operation_ids:
    - PostTerminalReadersReaderCollectInputs
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Collect inputs using a Reader

`POST /v1/terminal/readers/{reader}/collect_inputs`

Operation ID: `PostTerminalReadersReaderCollectInputs`

<p>Initiates an <a href="/docs/terminal/features/collect-inputs">input collection flow</a> on a Reader to display input forms and collect information from your customers.</p>

## Definition

```yaml
{"summary": "Collect inputs using a Reader", "description": "<p>Initiates an <a href=\"/docs/terminal/features/collect-inputs\">input collection flow</a> on a Reader to display input forms and collect information from your customers.</p>", "operationId": "PostTerminalReadersReaderCollectInputs", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["inputs"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "inputs": {"type": "array", "description": "List of inputs to be collected from the customer using the Reader. Maximum 5 inputs.", "items": {"title": "input_params", "required": ["custom_text", "type"], "type": "object", "properties": {"custom_text": {"title": "custom_text_params", "required": ["title"], "type": "object", "properties": {"description": {"maxLength": 500, "type": "string"}, "skip_button": {"maxLength": 14, "type": "string"}, "submit_button": {"maxLength": 30, "type": "string"}, "title": {"maxLength": 40, "type": "string"}}}, "required": {"type": "boolean"}, "selection": {"title": "selection_params", "required": ["choices"], "type": "object", "properties": {"choices": {"type": "array", "items": {"title": "choice_params", "required": ["id", "text"], "type": "object", "properties": {"id": {"maxLength": 50, "type": "string"}, "style": {"type": "string", "enum": ["primary", "secondary"]}, "text": {"maxLength": 30, "type": "string"}}}}}}, "toggles": {"type": "array", "items": {"title": "toggle_params", "type": "object", "properties": {"default_value": {"type": "string", "enum": ["disabled", "enabled"]}, "description": {"maxLength": 50, "type": "string"}, "title": {"maxLength": 50, "type": "string"}}}}, "type": {"type": "string", "enum": ["email", "numeric", "phone", "selection", "signature", "text"]}}}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "inputs": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
