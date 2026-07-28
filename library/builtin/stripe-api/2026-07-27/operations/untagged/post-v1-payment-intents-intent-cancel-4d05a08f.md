---
title: Cancel a PaymentIntent
page_id: operation-post-v1-payment-intents-intent-cancel-23921d02
path: operations/untagged
description: |-
    <p>You can cancel a PaymentIntent object when it’s in one of these statuses: <code>requires_payment_method</code>, <code>requires_capture</code>, <code>requires_confirmation</code>, <code>requires_action</code> or, <a href="/docs/payments/intents">in rare cases</a>, <code>processing</code>. </p>

    <p>After it’s canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a <code>status</code> of <code>requires_capture</code>, the remaining <code>amount_capturable</code> is automatically refunded. </p>

    <p>You can directly cancel the PaymentIntent for a Checkout Session only when the PaymentIntent has a status of <code>requires_capture</code>. Otherwise, you must <a href="/docs/api/checkout/sessions/expire">expire the Checkout Session</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_intents/{intent}/cancel
operation_ids:
    - PostPaymentIntentsIntentCancel
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Cancel a PaymentIntent

`POST /v1/payment_intents/{intent}/cancel`

Operation ID: `PostPaymentIntentsIntentCancel`

<p>You can cancel a PaymentIntent object when it’s in one of these statuses: <code>requires_payment_method</code>, <code>requires_capture</code>, <code>requires_confirmation</code>, <code>requires_action</code> or, <a href="/docs/payments/intents">in rare cases</a>, <code>processing</code>. </p>

<p>After it’s canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a <code>status</code> of <code>requires_capture</code>, the remaining <code>amount_capturable</code> is automatically refunded. </p>

<p>You can directly cancel the PaymentIntent for a Checkout Session only when the PaymentIntent has a status of <code>requires_capture</code>. Otherwise, you must <a href="/docs/api/checkout/sessions/expire">expire the Checkout Session</a>.</p>

## Definition

```yaml
{"summary": "Cancel a PaymentIntent", "description": "<p>You can cancel a PaymentIntent object when it’s in one of these statuses: <code>requires_payment_method</code>, <code>requires_capture</code>, <code>requires_confirmation</code>, <code>requires_action</code> or, <a href=\"/docs/payments/intents\">in rare cases</a>, <code>processing</code>. </p>\n\n<p>After it’s canceled, no additional charges are made by the PaymentIntent and any operations on the PaymentIntent fail with an error. For PaymentIntents with a <code>status</code> of <code>requires_capture</code>, the remaining <code>amount_capturable</code> is automatically refunded. </p>\n\n<p>You can directly cancel the PaymentIntent for a Checkout Session only when the PaymentIntent has a status of <code>requires_capture</code>. Otherwise, you must <a href=\"/docs/api/checkout/sessions/expire\">expire the Checkout Session</a>.</p>", "operationId": "PostPaymentIntentsIntentCancel", "parameters": [{"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"cancellation_reason": {"maxLength": 5000, "type": "string", "description": "Reason for canceling this PaymentIntent. Possible values are: `duplicate`, `fraudulent`, `requested_by_customer`, or `abandoned`", "enum": ["abandoned", "duplicate", "fraudulent", "requested_by_customer"], "x-stripeBypassValidation": true}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
