---
title: Attach a PaymentMethod to a Customer
page_id: operation-post-v1-payment-methods-payment-method-attach-d3255333
path: operations/untagged
description: |-
    <p>Attaches a PaymentMethod object to a Customer.</p>

    <p>To attach a new PaymentMethod to a customer for future payments, we recommend you use a <a href="/docs/api/setup_intents">SetupIntent</a>
    or a PaymentIntent with <a href="/docs/api/payment_intents/create#create_payment_intent-setup_future_usage">setup_future_usage</a>.
    These approaches will perform any necessary steps to set up the PaymentMethod for future payments. Using the <code>/v1/payment_methods/:id/attach</code>
    endpoint without first using a SetupIntent or PaymentIntent with <code>setup_future_usage</code> does not optimize the PaymentMethod for
    future use, which makes later declines and payment friction more likely.
    See <a href="/docs/payments/payment-intents#future-usage">Optimizing cards for future payments</a> for more information about setting up
    future payments.</p>

    <p>To use this PaymentMethod as the default for invoice or subscription payments,
    set <a href="/docs/api/customers/update#update_customer-invoice_settings-default_payment_method"><code>invoice_settings.default_payment_method</code></a>,
    on the Customer to the PaymentMethod’s ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_methods/{payment_method}/attach
operation_ids:
    - PostPaymentMethodsPaymentMethodAttach
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Attach a PaymentMethod to a Customer

`POST /v1/payment_methods/{payment_method}/attach`

Operation ID: `PostPaymentMethodsPaymentMethodAttach`

<p>Attaches a PaymentMethod object to a Customer.</p>

<p>To attach a new PaymentMethod to a customer for future payments, we recommend you use a <a href="/docs/api/setup_intents">SetupIntent</a>
or a PaymentIntent with <a href="/docs/api/payment_intents/create#create_payment_intent-setup_future_usage">setup_future_usage</a>.
These approaches will perform any necessary steps to set up the PaymentMethod for future payments. Using the <code>/v1/payment_methods/:id/attach</code>
endpoint without first using a SetupIntent or PaymentIntent with <code>setup_future_usage</code> does not optimize the PaymentMethod for
future use, which makes later declines and payment friction more likely.
See <a href="/docs/payments/payment-intents#future-usage">Optimizing cards for future payments</a> for more information about setting up
future payments.</p>

<p>To use this PaymentMethod as the default for invoice or subscription payments,
set <a href="/docs/api/customers/update#update_customer-invoice_settings-default_payment_method"><code>invoice_settings.default_payment_method</code></a>,
on the Customer to the PaymentMethod’s ID.</p>

## Definition

```yaml
{"summary": "Attach a PaymentMethod to a Customer", "description": "<p>Attaches a PaymentMethod object to a Customer.</p>\n\n<p>To attach a new PaymentMethod to a customer for future payments, we recommend you use a <a href=\"/docs/api/setup_intents\">SetupIntent</a>\nor a PaymentIntent with <a href=\"/docs/api/payment_intents/create#create_payment_intent-setup_future_usage\">setup_future_usage</a>.\nThese approaches will perform any necessary steps to set up the PaymentMethod for future payments. Using the <code>/v1/payment_methods/:id/attach</code>\nendpoint without first using a SetupIntent or PaymentIntent with <code>setup_future_usage</code> does not optimize the PaymentMethod for\nfuture use, which makes later declines and payment friction more likely.\nSee <a href=\"/docs/payments/payment-intents#future-usage\">Optimizing cards for future payments</a> for more information about setting up\nfuture payments.</p>\n\n<p>To use this PaymentMethod as the default for invoice or subscription payments,\nset <a href=\"/docs/api/customers/update#update_customer-invoice_settings-default_payment_method\"><code>invoice_settings.default_payment_method</code></a>,\non the Customer to the PaymentMethod’s ID.</p>", "operationId": "PostPaymentMethodsPaymentMethodAttach", "parameters": [{"name": "payment_method", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"customer": {"maxLength": 5000, "type": "string", "description": "The ID of the customer to which to attach the PaymentMethod."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The ID of the Account representing the customer to which to attach the PaymentMethod."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
