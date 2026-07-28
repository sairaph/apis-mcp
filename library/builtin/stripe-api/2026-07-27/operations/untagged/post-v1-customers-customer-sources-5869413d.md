---
title: Create a card
page_id: operation-post-v1-customers-customer-sources-933a6a64
path: operations/untagged
description: |-
    <p>When you create a new credit card, you must specify a customer or recipient on which to create it.</p>

    <p>If the card’s owner has no default card, then the new card will become the default.
    However, if the owner already has a default, then it will not change.
    To change the default, you should <a href="/api/customers/update">update the customer</a> to have a new <code>default_source</code>.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/customers/{customer}/sources
operation_ids:
    - PostCustomersCustomerSources
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a card

`POST /v1/customers/{customer}/sources`

Operation ID: `PostCustomersCustomerSources`

<p>When you create a new credit card, you must specify a customer or recipient on which to create it.</p>

<p>If the card’s owner has no default card, then the new card will become the default.
However, if the owner already has a default, then it will not change.
To change the default, you should <a href="/api/customers/update">update the customer</a> to have a new <code>default_source</code>.</p>

## Definition

```yaml
{"summary": "Create a card", "description": "<p>When you create a new credit card, you must specify a customer or recipient on which to create it.</p>\n\n<p>If the card’s owner has no default card, then the new card will become the default.\nHowever, if the owner already has a default, then it will not change.\nTo change the default, you should <a href=\"/api/customers/update\">update the customer</a> to have a new <code>default_source</code>.</p>", "operationId": "PostCustomersCustomerSources", "parameters": [{"name": "customer", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"alipay_account": {"maxLength": 5000, "type": "string", "description": "A token returned by [Stripe.js](https://stripe.com/docs/js) representing the user’s Alipay account details."}, "bank_account": {"description": "Either a token, like the ones returned by [Stripe.js](https://stripe.com/docs/js), or a dictionary containing a user's bank account details.", "anyOf": [{"title": "customer_payment_source_bank_account", "required": ["account_number", "country"], "type": "object", "properties": {"account_holder_name": {"maxLength": 5000, "type": "string"}, "account_holder_type": {"maxLength": 5000, "type": "string", "enum": ["company", "individual"]}, "account_number": {"maxLength": 5000, "type": "string"}, "country": {"maxLength": 5000, "type": "string"}, "currency": {"type": "string", "format": "currency"}, "object": {"maxLength": 5000, "type": "string", "enum": ["bank_account"]}, "routing_number": {"maxLength": 5000, "type": "string"}}}, {"maxLength": 5000, "type": "string"}]}, "card": {"description": "A token, like the ones returned by [Stripe.js](https://stripe.com/docs/js).", "anyOf": [{"title": "customer_payment_source_card", "required": ["exp_month", "exp_year", "number"], "type": "object", "properties": {"address_city": {"maxLength": 5000, "type": "string"}, "address_country": {"maxLength": 5000, "type": "string"}, "address_line1": {"maxLength": 5000, "type": "string"}, "address_line2": {"maxLength": 5000, "type": "string"}, "address_state": {"maxLength": 5000, "type": "string"}, "address_zip": {"maxLength": 5000, "type": "string"}, "cvc": {"maxLength": 5000, "type": "string"}, "encrypted": {"maxLength": 5000, "type": "string"}, "exp_month": {"type": "integer"}, "exp_year": {"type": "integer"}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}}, "name": {"maxLength": 5000, "type": "string"}, "network_token": {"title": "source_deprecated_card_network_token", "type": "object", "properties": {"number": {"maxLength": 5000, "type": "string"}}}, "number": {"maxLength": 5000, "type": "string"}, "object": {"maxLength": 5000, "type": "string", "enum": ["card"]}, "swipe_data": {"maxLength": 5000, "type": "string"}}}, {"maxLength": 5000, "type": "string"}], "x-stripeBypassValidation": true}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "source": {"maxLength": 5000, "type": "string", "description": "Please refer to full [documentation](https://api.stripe.com) instead.", "x-stripeBypassValidation": true}}, "additionalProperties": false}, "encoding": {"bank_account": {"style": "deepObject", "explode": true}, "card": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/payment_source"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
