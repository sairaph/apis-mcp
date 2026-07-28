---
title: Increment an authorization
page_id: operation-post-v1-payment-intents-intent-increment-authorization-a86f986d
path: operations/untagged
description: |-
    <p>Perform an incremental authorization on an eligible
    <a href="/docs/api/payment_intents/object">PaymentIntent</a>. To be eligible, the
    PaymentIntent’s status must be <code>requires_capture</code> and
    <a href="/docs/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported">incremental_authorization_supported</a>
    must be <code>true</code>.</p>

    <p>Incremental authorizations attempt to increase the authorized amount on
    your customer’s card to the new, higher <code>amount</code> provided. Similar to the
    initial authorization, incremental authorizations can be declined. A
    single PaymentIntent can call this endpoint multiple times to further
    increase the authorized amount.</p>

    <p>If the incremental authorization succeeds, the PaymentIntent object
    returns with the updated
    <a href="/docs/api/payment_intents/object#payment_intent_object-amount">amount</a>.
    If the incremental authorization fails, a
    <a href="/docs/error-codes#card-declined">card_declined</a> error returns, and no other
    fields on the PaymentIntent or Charge update. The PaymentIntent
    object remains capturable for the previously authorized amount.</p>

    <p>Each PaymentIntent can have a maximum of 10 incremental authorization attempts, including declines.
    After it’s captured, a PaymentIntent can no longer be incremented.</p>

    <p>Learn more about incremental authorizations with
    <a href="/docs/terminal/features/incremental-authorizations">in-person payments</a> and
    <a href="/docs/payments/incremental-authorization?platform=web&ui=elements">online payments</a>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_intents/{intent}/increment_authorization
operation_ids:
    - PostPaymentIntentsIntentIncrementAuthorization
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Increment an authorization

`POST /v1/payment_intents/{intent}/increment_authorization`

Operation ID: `PostPaymentIntentsIntentIncrementAuthorization`

<p>Perform an incremental authorization on an eligible
<a href="/docs/api/payment_intents/object">PaymentIntent</a>. To be eligible, the
PaymentIntent’s status must be <code>requires_capture</code> and
<a href="/docs/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported">incremental_authorization_supported</a>
must be <code>true</code>.</p>

<p>Incremental authorizations attempt to increase the authorized amount on
your customer’s card to the new, higher <code>amount</code> provided. Similar to the
initial authorization, incremental authorizations can be declined. A
single PaymentIntent can call this endpoint multiple times to further
increase the authorized amount.</p>

<p>If the incremental authorization succeeds, the PaymentIntent object
returns with the updated
<a href="/docs/api/payment_intents/object#payment_intent_object-amount">amount</a>.
If the incremental authorization fails, a
<a href="/docs/error-codes#card-declined">card_declined</a> error returns, and no other
fields on the PaymentIntent or Charge update. The PaymentIntent
object remains capturable for the previously authorized amount.</p>

<p>Each PaymentIntent can have a maximum of 10 incremental authorization attempts, including declines.
After it’s captured, a PaymentIntent can no longer be incremented.</p>

<p>Learn more about incremental authorizations with
<a href="/docs/terminal/features/incremental-authorizations">in-person payments</a> and
<a href="/docs/payments/incremental-authorization?platform=web&ui=elements">online payments</a>.</p>

## Definition

```yaml
{"summary": "Increment an authorization", "description": "<p>Perform an incremental authorization on an eligible\n<a href=\"/docs/api/payment_intents/object\">PaymentIntent</a>. To be eligible, the\nPaymentIntent’s status must be <code>requires_capture</code> and\n<a href=\"/docs/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported\">incremental_authorization_supported</a>\nmust be <code>true</code>.</p>\n\n<p>Incremental authorizations attempt to increase the authorized amount on\nyour customer’s card to the new, higher <code>amount</code> provided. Similar to the\ninitial authorization, incremental authorizations can be declined. A\nsingle PaymentIntent can call this endpoint multiple times to further\nincrease the authorized amount.</p>\n\n<p>If the incremental authorization succeeds, the PaymentIntent object\nreturns with the updated\n<a href=\"/docs/api/payment_intents/object#payment_intent_object-amount\">amount</a>.\nIf the incremental authorization fails, a\n<a href=\"/docs/error-codes#card-declined\">card_declined</a> error returns, and no other\nfields on the PaymentIntent or Charge update. The PaymentIntent\nobject remains capturable for the previously authorized amount.</p>\n\n<p>Each PaymentIntent can have a maximum of 10 incremental authorization attempts, including declines.\nAfter it’s captured, a PaymentIntent can no longer be incremented.</p>\n\n<p>Learn more about incremental authorizations with\n<a href=\"/docs/terminal/features/incremental-authorizations\">in-person payments</a> and\n<a href=\"/docs/payments/incremental-authorization?platform=web&ui=elements\">online payments</a>.</p>", "operationId": "PostPaymentIntentsIntentIncrementAuthorization", "parameters": [{"name": "intent", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["amount"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The updated total amount that you intend to collect from the cardholder. This amount must be greater than the currently authorized amount."}, "amount_details": {"title": "amount_details_param", "type": "object", "properties": {"discount_amount": {"anyOf": [{"type": "integer"}, {"type": "string", "enum": [""]}]}, "enforce_arithmetic_validation": {"type": "boolean"}, "line_items": {"anyOf": [{"type": "array", "items": {"title": "amount_details_line_item_param", "required": ["product_name", "quantity", "unit_cost"], "type": "object", "properties": {"discount_amount": {"type": "integer"}, "payment_method_options": {"title": "amount_details_line_item_payment_method_options_param", "type": "object", "properties": {"card": {"title": "payment_intent_amount_details_line_item_payment_method_options_param", "type": "object", "properties": {"commodity_code": {"maxLength": 12, "type": "string"}}}, "card_present": {"title": "amount_details_line_item_payment_method_options_param", "type": "object", "properties": {"commodity_code": {"maxLength": 12, "type": "string"}}}, "klarna": {"title": "payment_intent_amount_details_line_item_payment_method_options_param", "type": "object", "properties": {"image_url": {"maxLength": 4096, "type": "string"}, "product_url": {"maxLength": 4096, "type": "string"}, "reference": {"maxLength": 255, "type": "string"}, "subscription_reference": {"maxLength": 255, "type": "string"}}}, "paypal": {"title": "amount_details_line_item_payment_method_options_param", "type": "object", "properties": {"category": {"type": "string", "enum": ["digital_goods", "donation", "physical_goods"]}, "description": {"maxLength": 127, "type": "string"}, "sold_by": {"maxLength": 127, "type": "string"}}}}}, "product_code": {"maxLength": 12, "type": "string"}, "product_name": {"maxLength": 1024, "type": "string"}, "quantity": {"type": "integer"}, "tax": {"title": "amount_details_line_item_tax_param", "required": ["total_tax_amount"], "type": "object", "properties": {"total_tax_amount": {"type": "integer"}}}, "unit_cost": {"type": "integer"}, "unit_of_measure": {"maxLength": 12, "type": "string"}}}}, {"type": "string", "enum": [""]}]}, "shipping": {"anyOf": [{"title": "amount_details_shipping_param", "type": "object", "properties": {"amount": {"anyOf": [{"type": "integer"}, {"type": "string", "enum": [""]}]}, "from_postal_code": {"anyOf": [{"maxLength": 10, "type": "string"}, {"type": "string", "enum": [""]}]}, "to_postal_code": {"anyOf": [{"maxLength": 10, "type": "string"}, {"type": "string", "enum": [""]}]}}}, {"type": "string", "enum": [""]}]}, "tax": {"anyOf": [{"title": "amount_details_tax_param", "required": ["total_tax_amount"], "type": "object", "properties": {"total_tax_amount": {"type": "integer"}}}, {"type": "string", "enum": [""]}]}}, "description": "Provides industry-specific information about the amount."}, "application_fee_amount": {"type": "integer", "description": "The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total amount captured. For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts)."}, "description": {"maxLength": 1000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "hooks": {"title": "async_workflows_param", "type": "object", "properties": {"inputs": {"title": "async_workflows_inputs_param", "type": "object", "properties": {"tax": {"title": "async_workflows_inputs_tax_param", "required": ["calculation"], "type": "object", "properties": {"calculation": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}}}}}}, "description": "Automations to be run during the PaymentIntent lifecycle"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "payment_details": {"title": "payment_details_order_customer_reference_param", "type": "object", "properties": {"customer_reference": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "order_reference": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}}, "description": "Provides industry-specific information about the charge."}, "statement_descriptor": {"maxLength": 22, "type": "string", "description": "Text that appears on the customer's statement as the statement descriptor for a non-card or card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors)."}, "transfer_data": {"title": "transfer_data_update_auth_params", "type": "object", "properties": {"amount": {"type": "integer"}}, "description": "The parameters used to automatically create a transfer after the payment is captured.\nLearn more about the [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts)."}}, "additionalProperties": false}, "encoding": {"amount_details": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "hooks": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payment_details": {"style": "deepObject", "explode": true}, "transfer_data": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_intent"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
