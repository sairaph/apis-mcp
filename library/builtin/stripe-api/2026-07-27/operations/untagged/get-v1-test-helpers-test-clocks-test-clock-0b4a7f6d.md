---
title: Retrieve a test clock
page_id: operation-get-v1-test-helpers-test-clocks-test-clock-a0cdb420
path: operations/untagged
description: <p>Retrieves a test clock.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/test_helpers/test_clocks/{test_clock}
operation_ids:
    - GetTestHelpersTestClocksTestClock
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a test clock

`GET /v1/test_helpers/test_clocks/{test_clock}`

Operation ID: `GetTestHelpersTestClocksTestClock`

<p>Retrieves a test clock.</p>

## Definition

```yaml
{"summary": "Retrieve a test clock", "description": "<p>Retrieves a test clock.</p>", "operationId": "GetTestHelpersTestClocksTestClock", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "test_clock", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/test_helpers.test_clock"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
