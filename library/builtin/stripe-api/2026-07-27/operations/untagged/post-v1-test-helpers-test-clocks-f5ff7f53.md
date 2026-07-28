---
title: Create a test clock
page_id: operation-post-v1-test-helpers-test-clocks-2c877185
path: operations/untagged
description: <p>Creates a new test clock that can be attached to new customers and quotes.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/test_clocks
operation_ids:
    - PostTestHelpersTestClocks
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a test clock

`POST /v1/test_helpers/test_clocks`

Operation ID: `PostTestHelpersTestClocks`

<p>Creates a new test clock that can be attached to new customers and quotes.</p>

## Definition

```yaml
{"summary": "Create a test clock", "description": "<p>Creates a new test clock that can be attached to new customers and quotes.</p>", "operationId": "PostTestHelpersTestClocks", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["frozen_time"], "type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "Existing customer this test clock will be attached to. Once attached, customers can't be removed from a test clock."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "frozen_time": {"type": "integer", "description": "The initial frozen time for this test clock.", "format": "unix-time"}, "name": {"maxLength": 300, "type": "string", "description": "The name for this test clock."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/test_helpers.test_clock"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
