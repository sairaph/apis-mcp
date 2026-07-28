---
title: Expire a Checkout Session
page_id: operation-post-v1-checkout-sessions-session-expire-f7657d86
path: operations/untagged
description: |-
    <p>A Checkout Session can be expired when it is in one of these statuses: <code>open</code> </p>

    <p>After it expires, a customer can’t complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/checkout/sessions/{session}/expire
operation_ids:
    - PostCheckoutSessionsSessionExpire
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Expire a Checkout Session

`POST /v1/checkout/sessions/{session}/expire`

Operation ID: `PostCheckoutSessionsSessionExpire`

<p>A Checkout Session can be expired when it is in one of these statuses: <code>open</code> </p>

<p>After it expires, a customer can’t complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.</p>

## Definition

```yaml
{"summary": "Expire a Checkout Session", "description": "<p>A Checkout Session can be expired when it is in one of these statuses: <code>open</code> </p>\n\n<p>After it expires, a customer can’t complete a Checkout Session and customers loading the Checkout Session see a message saying the Checkout Session is expired.</p>", "operationId": "PostCheckoutSessionsSessionExpire", "parameters": [{"name": "session", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/checkout.session"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
