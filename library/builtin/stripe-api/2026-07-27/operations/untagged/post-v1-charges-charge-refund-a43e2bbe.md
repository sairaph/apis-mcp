---
title: Create a refund
page_id: operation-post-v1-charges-charge-refund-32c9cdf7
path: operations/untagged
description: |-
    <p>When you create a new refund, you must specify either a Charge or a PaymentIntent object.</p>

    <p>This action refunds a previously created charge that’s not refunded yet.
    Funds are refunded to the credit or debit card that’s originally charged.</p>

    <p>You can optionally refund only part of a charge.
    You can repeat this until the entire charge is refunded.</p>

    <p>After you entirely refund a charge, you can’t refund it again.
    This method raises an error when it’s called on an already-refunded charge,
    or when you attempt to refund more money than is left on a charge.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/charges/{charge}/refund
operation_ids:
    - PostChargesChargeRefund
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a refund

`POST /v1/charges/{charge}/refund`

Operation ID: `PostChargesChargeRefund`

<p>When you create a new refund, you must specify either a Charge or a PaymentIntent object.</p>

<p>This action refunds a previously created charge that’s not refunded yet.
Funds are refunded to the credit or debit card that’s originally charged.</p>

<p>You can optionally refund only part of a charge.
You can repeat this until the entire charge is refunded.</p>

<p>After you entirely refund a charge, you can’t refund it again.
This method raises an error when it’s called on an already-refunded charge,
or when you attempt to refund more money than is left on a charge.</p>

## Definition

```yaml
{"summary": "Create a refund", "description": "<p>When you create a new refund, you must specify either a Charge or a PaymentIntent object.</p>\n\n<p>This action refunds a previously created charge that’s not refunded yet.\nFunds are refunded to the credit or debit card that’s originally charged.</p>\n\n<p>You can optionally refund only part of a charge.\nYou can repeat this until the entire charge is refunded.</p>\n\n<p>After you entirely refund a charge, you can’t refund it again.\nThis method raises an error when it’s called on an already-refunded charge,\nor when you attempt to refund more money than is left on a charge.</p>", "operationId": "PostChargesChargeRefund", "parameters": [{"name": "charge", "in": "path", "description": "The identifier of the charge to refund.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) representing how much of this charge to refund. Can refund only up to the remaining, unrefunded amount of the charge."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "instructions_email": {"type": "string", "description": "For payment methods without native refund support (e.g., Konbini, PromptPay), use this email from the customer to receive refund instructions."}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The identifier of the PaymentIntent to refund."}, "reason": {"maxLength": 5000, "type": "string", "description": "String indicating the reason for the refund. If set, possible values are `duplicate`, `fraudulent`, and `requested_by_customer`. If you believe the charge to be fraudulent, specifying `fraudulent` as the reason will add the associated card and email to your [block lists](https://docs.stripe.com/radar/lists), and will also help us improve our fraud detection algorithms.", "enum": ["duplicate", "fraudulent", "requested_by_customer"]}, "refund_application_fee": {"type": "boolean", "description": "Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge."}, "reverse_transfer": {"type": "boolean", "description": "Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount).<br><br>A transfer can be reversed only by the application that created the charge."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/charge"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
