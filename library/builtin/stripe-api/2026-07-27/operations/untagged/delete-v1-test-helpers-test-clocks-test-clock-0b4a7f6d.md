---
title: Delete a test clock
page_id: operation-delete-v1-test-helpers-test-clocks-test-clock-911ac4d7
path: operations/untagged
description: <p>Deletes a test clock.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/test_helpers/test_clocks/{test_clock}
operation_ids:
    - DeleteTestHelpersTestClocksTestClock
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a test clock

`DELETE /v1/test_helpers/test_clocks/{test_clock}`

Operation ID: `DeleteTestHelpersTestClocksTestClock`

<p>Deletes a test clock.</p>

## Definition

```yaml
{"summary": "Delete a test clock", "description": "<p>Deletes a test clock.</p>", "operationId": "DeleteTestHelpersTestClocksTestClock", "parameters": [{"name": "test_clock", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_test_helpers.test_clock"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
