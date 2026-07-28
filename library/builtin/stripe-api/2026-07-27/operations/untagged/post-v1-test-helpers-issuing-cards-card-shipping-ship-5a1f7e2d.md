---
title: Ship a testmode card
page_id: operation-post-v1-test-helpers-issuing-cards-card-shipping-ship-b8438d42
path: operations/untagged
description: <p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>shipped</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/test_helpers/issuing/cards/{card}/shipping/ship
operation_ids:
    - PostTestHelpersIssuingCardsCardShippingShip
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Ship a testmode card

`POST /v1/test_helpers/issuing/cards/{card}/shipping/ship`

Operation ID: `PostTestHelpersIssuingCardsCardShippingShip`

<p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>shipped</code>.</p>

## Definition

```yaml
{"summary": "Ship a testmode card", "description": "<p>Updates the shipping status of the specified Issuing <code>Card</code> object to <code>shipped</code>.</p>", "operationId": "PostTestHelpersIssuingCardsCardShippingShip", "parameters": [{"name": "card", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/issuing.card"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
