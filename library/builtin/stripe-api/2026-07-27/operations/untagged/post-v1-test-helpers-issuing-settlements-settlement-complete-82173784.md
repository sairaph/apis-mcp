---
title: Complete a test-mode settlement
page_id: operation-post-v1-test-helpers-issuing-settlements-settlement-complete-e5cb46fb
path: operations/untagged
description: <p>Allows the user to mark an Issuing settlement as complete.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/settlements/{settlement}/complete
operation_ids:
    - PostTestHelpersIssuingSettlementsSettlementComplete
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Complete a test-mode settlement

`POST /v1/test_helpers/issuing/settlements/{settlement}/complete`

Operation ID: `PostTestHelpersIssuingSettlementsSettlementComplete`

<p>Allows the user to mark an Issuing settlement as complete.</p>

## Definition

```yaml
{"summary": "Complete a test-mode settlement", "description": "<p>Allows the user to mark an Issuing settlement as complete.</p>", "operationId": "PostTestHelpersIssuingSettlementsSettlementComplete", "parameters": [{"name": "settlement", "in": "path", "description": "The settlement token to mark as complete.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.settlement"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
