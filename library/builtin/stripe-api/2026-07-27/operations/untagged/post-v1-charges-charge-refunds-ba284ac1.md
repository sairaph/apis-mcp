---
title: Create customer balance refund
page_id: operation-post-v1-charges-charge-refunds-eb762093
path: operations/untagged
description: |-
    <p>When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.</p>

    <p>Creating a new refund will refund a charge that has previously been created but not yet refunded.
    Funds will be refunded to the credit or debit card that was originally charged.</p>

    <p>You can optionally refund only part of a charge.
    You can do so multiple times, until the entire charge has been refunded.</p>

    <p>Once entirely refunded, a charge can’t be refunded again.
    This method will raise an error when called on an already-refunded charge,
    or when trying to refund more money than is left on a charge.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}/refunds
operation_ids:
    - PostChargesChargeRefunds
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create customer balance refund

`POST /v1/charges/{charge}/refunds`

Operation ID: `PostChargesChargeRefunds`

<p>When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.</p>

<p>Creating a new refund will refund a charge that has previously been created but not yet refunded.
Funds will be refunded to the credit or debit card that was originally charged.</p>

<p>You can optionally refund only part of a charge.
You can do so multiple times, until the entire charge has been refunded.</p>

<p>Once entirely refunded, a charge can’t be refunded again.
This method will raise an error when called on an already-refunded charge,
or when trying to refund more money than is left on a charge.</p>

## Definition

```yaml
{"summary": "Create customer balance refund", "description": "<p>When you create a new refund, you must specify a Charge or a PaymentIntent object on which to create it.</p>\n\n<p>Creating a new refund will refund a charge that has previously been created but not yet refunded.\nFunds will be refunded to the credit or debit card that was originally charged.</p>\n\n<p>You can optionally refund only part of a charge.\nYou can do so multiple times, until the entire charge has been refunded.</p>\n\n<p>Once entirely refunded, a charge can’t be refunded again.\nThis method will raise an error when called on an already-refunded charge,\nor when trying to refund more money than is left on a charge.</p>", "operationId": "PostChargesChargeRefunds", "parameters": [{"name": "charge", "in": "path", "description": "The identifier of the charge to refund.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer"}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "customer": {"maxLength": 5000, "type": "string", "description": "Customer whose customer balance to refund from."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "instructions_email": {"type": "string", "description": "For payment methods without native refund support (e.g., Konbini, PromptPay), use this email from the customer to receive refund instructions."}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "origin": {"type": "string", "description": "Origin of the refund", "enum": ["customer_balance"]}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The identifier of the PaymentIntent to refund."}, "reason": {"maxLength": 5000, "type": "string", "description": "String indicating the reason for the refund. If set, possible values are `duplicate`, `fraudulent`, and `requested_by_customer`. If you believe the charge to be fraudulent, specifying `fraudulent` as the reason will add the associated card and email to your [block lists](https://docs.stripe.com/radar/lists), and will also help us improve our fraud detection algorithms.", "enum": ["duplicate", "fraudulent", "requested_by_customer"]}, "refund_application_fee": {"type": "boolean", "description": "Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge."}, "reverse_transfer": {"type": "boolean", "description": "Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount).<br><br>A transfer can be reversed only by the application that created the charge."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/refund"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
