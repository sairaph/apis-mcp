---
title: Reconcile a customer_balance PaymentIntent
page_id: operation-post-v1-payment-intents-intent-apply-customer-balance-ec6fba1a
path: operations/untagged
description: <p>Manually reconcile the remaining amount for a <code>customer_balance</code> PaymentIntent.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_intents/{intent}/apply_customer_balance
operation_ids:
    - PostPaymentIntentsIntentApplyCustomerBalance
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Reconcile a customer_balance PaymentIntent

`POST /v1/payment_intents/{intent}/apply_customer_balance`

Operation ID: `PostPaymentIntentsIntentApplyCustomerBalance`

<p>Manually reconcile the remaining amount for a <code>customer_balance</code> PaymentIntent.</p>

## Definition

```yaml
{"summary": "Reconcile a customer_balance PaymentIntent", "description": "<p>Manually reconcile the remaining amount for a <code>customer_balance</code> PaymentIntent.</p>", "operationId": "PostPaymentIntentsIntentApplyCustomerBalance", "parameters": [{"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "Amount that you intend to apply to this PaymentIntent from the customer’s cash balance. If the PaymentIntent was created by an Invoice, the full amount of the PaymentIntent is applied regardless of this parameter.\n\nA positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1 USD or 100 to charge 100 JPY, a zero-decimal currency). The maximum amount is the amount of the PaymentIntent.\n\nWhen you omit the amount, it defaults to the remaining amount requested on the PaymentIntent."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
