---
title: Update a PaymentMethod
page_id: operation-post-v1-payment-methods-payment-method-c272fedb
path: operations/untagged
description: <p>Updates a PaymentMethod object. A PaymentMethod must be attached to a customer to be updated.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/payment_methods/{payment_method}
operation_ids:
    - PostPaymentMethodsPaymentMethod
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a PaymentMethod

`POST /v1/payment_methods/{payment_method}`

Operation ID: `PostPaymentMethodsPaymentMethod`

<p>Updates a PaymentMethod object. A PaymentMethod must be attached to a customer to be updated.</p>

## Definition

```yaml
{"summary": "Update a PaymentMethod", "description": "<p>Updates a PaymentMethod object. A PaymentMethod must be attached to a customer to be updated.</p>", "operationId": "PostPaymentMethodsPaymentMethod", "parameters": [{"name": "payment_method", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"allow_redisplay": {"type": "string", "description": "This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.", "enum": ["always", "limited", "unspecified"]}, "billing_details": {"title": "billing_details_inner_params", "type": "object", "properties": {"address": {"anyOf": [{"title": "billing_details_address", "type": "object", "properties": {"city": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "line1": {"maxLength": 5000, "type": "string"}, "line2": {"maxLength": 5000, "type": "string"}, "postal_code": {"maxLength": 5000, "type": "string"}, "state": {"maxLength": 5000, "type": "string"}}}, {"type": "string", "enum": [""]}]}, "email": {"anyOf": [{"type": "string"}, {"type": "string", "enum": [""]}]}, "name": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "phone": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "tax_id": {"maxLength": 5000, "type": "string"}}, "description": "Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods."}, "card": {"title": "update_api_param", "type": "object", "properties": {"exp_month": {"type": "integer"}, "exp_year": {"type": "integer"}, "networks": {"title": "networks_update_api_param", "type": "object", "properties": {"preferred": {"type": "string", "enum": ["", "cartes_bancaires", "mastercard", "visa"]}}}}, "description": "If this is a `card` PaymentMethod, this hash contains the user's card details."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "payto": {"title": "param", "type": "object", "properties": {"account_number": {"maxLength": 5000, "type": "string"}, "bsb_number": {"maxLength": 5000, "type": "string"}, "pay_id": {"maxLength": 5000, "type": "string"}}, "description": "If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method."}, "us_bank_account": {"title": "update_param", "type": "object", "properties": {"account_holder_type": {"type": "string", "enum": ["company", "individual"]}, "account_type": {"type": "string", "enum": ["checking", "savings"]}}, "description": "If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method."}}, "additionalProperties": false}, "encoding": {"billing_details": {"style": "deepObject", "explode": true}, "card": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "payto": {"style": "deepObject", "explode": true}, "us_bank_account": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_method"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
