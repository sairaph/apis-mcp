---
title: Update a price
page_id: operation-post-v1-prices-price-3932f91a
path: operations/untagged
description: <p>Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/prices/{price}
operation_ids:
    - PostPricesPrice
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a price

`POST /v1/prices/{price}`

Operation ID: `PostPricesPrice`

<p>Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.</p>

## Definition

```yaml
{"summary": "Update a price", "description": "<p>Updates the specified price by setting the values of the parameters passed. Any parameters not provided are left unchanged.</p>", "operationId": "PostPricesPrice", "parameters": [{"name": "price", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the price can be used for new purchases. Defaults to `true`."}, "currency_options": {"description": "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).", "anyOf": [{"type": "object", "additionalProperties": {"title": "currency_option", "type": "object", "properties": {"custom_unit_amount": {"title": "custom_unit_amount", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean"}, "maximum": {"type": "integer"}, "minimum": {"type": "integer"}, "preset": {"type": "integer"}}}, "tax_behavior": {"type": "string", "enum": ["exclusive", "inclusive", "unspecified"]}, "tiers": {"type": "array", "items": {"title": "tier", "required": ["up_to"], "type": "object", "properties": {"flat_amount": {"type": "integer"}, "flat_amount_decimal": {"type": "string", "format": "decimal"}, "unit_amount": {"type": "integer"}, "unit_amount_decimal": {"type": "string", "format": "decimal"}, "up_to": {"anyOf": [{"maxLength": 5000, "type": "string", "enum": ["inf"]}, {"type": "integer"}]}}}}, "unit_amount": {"type": "integer"}, "unit_amount_decimal": {"type": "string", "format": "decimal"}}}}, {"type": "string", "enum": [""]}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "lookup_key": {"maxLength": 200, "type": "string", "description": "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters."}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "nickname": {"maxLength": 5000, "type": "string", "description": "A brief description of the price, hidden from customers."}, "tax_behavior": {"type": "string", "description": "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.", "enum": ["exclusive", "inclusive", "unspecified"]}, "transfer_lookup_key": {"type": "boolean", "description": "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price."}}, "additionalProperties": false}, "encoding": {"currency_options": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/price"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
