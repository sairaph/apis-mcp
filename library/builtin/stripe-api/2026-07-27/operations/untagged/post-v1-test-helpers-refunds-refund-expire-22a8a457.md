---
title: Expire a pending refund.
page_id: operation-post-v1-test-helpers-refunds-refund-expire-305c1563
path: operations/untagged
description: <p>Expire a refund with a status of <code>requires_action</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/refunds/{refund}/expire
operation_ids:
    - PostTestHelpersRefundsRefundExpire
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Expire a pending refund.

`POST /v1/test_helpers/refunds/{refund}/expire`

Operation ID: `PostTestHelpersRefundsRefundExpire`

<p>Expire a refund with a status of <code>requires_action</code>.</p>

## Definition

```yaml
{"summary": "Expire a pending refund.", "description": "<p>Expire a refund with a status of <code>requires_action</code>.</p>", "operationId": "PostTestHelpersRefundsRefundExpire", "parameters": [{"name": "refund", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
