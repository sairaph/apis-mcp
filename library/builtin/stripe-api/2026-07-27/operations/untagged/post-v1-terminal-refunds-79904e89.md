---
title: Create a refund using a Terminal-supported device.
page_id: operation-post-v1-terminal-refunds-28f31a3a
path: operations/untagged
description: |-
    <p>Internal endpoint for terminal use to create a refund for a card_present or card charge.</p>

    <p>You can optionally refund only part of a charge.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/terminal/refunds
operation_ids:
    - PostTerminalRefunds
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a refund using a Terminal-supported device.

`POST /v1/terminal/refunds`

Operation ID: `PostTerminalRefunds`

<p>Internal endpoint for terminal use to create a refund for a card_present or card charge.</p>

<p>You can optionally refund only part of a charge.</p>

## Definition

```yaml
{"summary": "Create a refund using a Terminal-supported device.", "description": "<p>Internal endpoint for terminal use to create a refund for a card_present or card charge.</p>\n\n<p>You can optionally refund only part of a charge.</p>", "operationId": "PostTerminalRefunds", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/terminal.refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
