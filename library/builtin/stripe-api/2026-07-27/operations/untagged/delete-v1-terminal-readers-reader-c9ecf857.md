---
title: Delete a Reader
page_id: operation-delete-v1-terminal-readers-reader-59fd475d
path: operations/untagged
description: <p>Deletes a <code>Reader</code> object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/terminal/readers/{reader}
operation_ids:
    - DeleteTerminalReadersReader
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a Reader

`DELETE /v1/terminal/readers/{reader}`

Operation ID: `DeleteTerminalReadersReader`

<p>Deletes a <code>Reader</code> object.</p>

## Definition

```yaml
{"summary": "Delete a Reader", "description": "<p>Deletes a <code>Reader</code> object.</p>", "operationId": "DeleteTerminalReadersReader", "parameters": [{"name": "reader", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_terminal.reader"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
