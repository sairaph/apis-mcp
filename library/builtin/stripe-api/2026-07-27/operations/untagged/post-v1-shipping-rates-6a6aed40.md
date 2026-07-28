---
title: Create a shipping rate
page_id: operation-post-v1-shipping-rates-af2cea76
path: operations/untagged
description: <p>Creates a new shipping rate object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/shipping_rates
operation_ids:
    - PostShippingRates
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a shipping rate

`POST /v1/shipping_rates`

Operation ID: `PostShippingRates`

<p>Creates a new shipping rate object.</p>

## Definition

```yaml
{"summary": "Create a shipping rate", "description": "<p>Creates a new shipping rate object.</p>", "operationId": "PostShippingRates", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["display_name"], "type": "object", "properties": {"delivery_estimate": {"title": "delivery_estimate", "type": "object", "properties": {"maximum": {"title": "delivery_estimate_bound", "required": ["unit", "value"], "type": "object", "properties": {"unit": {"type": "string", "enum": ["business_day", "day", "hour", "month", "week"]}, "value": {"type": "integer"}}}, "minimum": {"title": "delivery_estimate_bound", "required": ["unit", "value"], "type": "object", "properties": {"unit": {"type": "string", "enum": ["business_day", "day", "hour", "month", "week"]}, "value": {"type": "integer"}}}}, "description": "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions."}, "display_name": {"maxLength": 100, "type": "string", "description": "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "fixed_amount": {"title": "fixed_amount", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer"}, "currency": {"type": "string", "format": "currency"}, "currency_options": {"type": "object", "additionalProperties": {"title": "currency_option", "required": ["amount"], "type": "object", "properties": {"amount": {"type": "integer"}, "tax_behavior": {"type": "string", "enum": ["exclusive", "inclusive", "unspecified"]}}}}}, "description": "Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "tax_behavior": {"type": "string", "description": "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.", "enum": ["exclusive", "inclusive", "unspecified"]}, "tax_code": {"type": "string", "description": "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`."}, "type": {"type": "string", "description": "The type of calculation to use on the shipping rate.", "enum": ["fixed_amount"]}}, "additionalProperties": false}, "encoding": {"delivery_estimate": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "fixed_amount": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/shipping_rate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
