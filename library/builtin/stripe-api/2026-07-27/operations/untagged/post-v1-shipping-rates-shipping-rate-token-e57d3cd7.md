---
title: Update a shipping rate
page_id: operation-post-v1-shipping-rates-shipping-rate-token-7f7ea058
path: operations/untagged
description: <p>Updates an existing shipping rate object.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/shipping_rates/{shipping_rate_token}
operation_ids:
    - PostShippingRatesShippingRateToken
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a shipping rate

`POST /v1/shipping_rates/{shipping_rate_token}`

Operation ID: `PostShippingRatesShippingRateToken`

<p>Updates an existing shipping rate object.</p>

## Definition

```yaml
{"summary": "Update a shipping rate", "description": "<p>Updates an existing shipping rate object.</p>", "operationId": "PostShippingRatesShippingRateToken", "parameters": [{"name": "shipping_rate_token", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the shipping rate can be used for new purchases. Defaults to `true`."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "fixed_amount": {"title": "fixed_amount_update", "type": "object", "properties": {"currency_options": {"type": "object", "additionalProperties": {"title": "currency_option_update", "type": "object", "properties": {"amount": {"type": "integer"}, "tax_behavior": {"type": "string", "enum": ["exclusive", "inclusive", "unspecified"]}}}}}, "description": "Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`."}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "tax_behavior": {"type": "string", "description": "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.", "enum": ["exclusive", "inclusive", "unspecified"]}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "fixed_amount": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/shipping_rate"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
