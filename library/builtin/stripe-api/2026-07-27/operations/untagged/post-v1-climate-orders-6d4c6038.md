---
title: Create an order
page_id: operation-post-v1-climate-orders-c219f2a8
path: operations/untagged
description: |-
    <p>Creates a Climate order object for a given Climate product. The order will be processed immediately
    after creation and payment will be deducted your Stripe balance.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/climate/orders
operation_ids:
    - PostClimateOrders
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create an order

`POST /v1/climate/orders`

Operation ID: `PostClimateOrders`

<p>Creates a Climate order object for a given Climate product. The order will be processed immediately
after creation and payment will be deducted your Stripe balance.</p>

## Definition

```yaml
{"summary": "Create an order", "description": "<p>Creates a Climate order object for a given Climate product. The order will be processed immediately\nafter creation and payment will be deducted your Stripe balance.</p>", "operationId": "PostClimateOrders", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["product"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Requested amount of carbon removal units. Either this or `metric_tons` must be specified."}, "beneficiary": {"title": "beneficiary_params", "required": ["public_name"], "type": "object", "properties": {"public_name": {"maxLength": 5000, "type": "string"}}, "description": "Publicly sharable reference for the end beneficiary of carbon removal. Assumed to be the Stripe account if not set."}, "currency": {"maxLength": 5000, "type": "string", "description": "Request currency for the order as a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a supported [settlement currency for your account](https://stripe.com/docs/currencies). If omitted, the account's default currency will be used."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "metric_tons": {"type": "string", "description": "Requested number of tons for the order. Either this or `amount` must be specified.", "format": "decimal"}, "product": {"maxLength": 5000, "type": "string", "description": "Unique identifier of the Climate product."}}, "additionalProperties": false}, "encoding": {"beneficiary": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/climate.order"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
