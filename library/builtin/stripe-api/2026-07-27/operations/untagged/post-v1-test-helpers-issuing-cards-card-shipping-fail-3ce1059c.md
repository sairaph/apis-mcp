---
title: Fail a testmode card
page_id: operation-post-v1-test-helpers-issuing-cards-card-shipping-fail-75a31466
path: operations/untagged
description: <p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>failure</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/cards/{card}/shipping/fail
operation_ids:
    - PostTestHelpersIssuingCardsCardShippingFail
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Fail a testmode card

`POST /v1/test_helpers/issuing/cards/{card}/shipping/fail`

Operation ID: `PostTestHelpersIssuingCardsCardShippingFail`

<p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>failure</code>.</p>

## Definition

```yaml
{"summary": "Fail a testmode card", "description": "<p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>failure</code>.</p>", "operationId": "PostTestHelpersIssuingCardsCardShippingFail", "parameters": [{"name": "card", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.card"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
