---
title: Update an order
page_id: operation-post-v1-climate-orders-order-8acf2aba
path: operations/untagged
description: <p>Updates the specified order by setting the values of the parameters passed.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/climate/orders/{order}
operation_ids:
    - PostClimateOrdersOrder
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update an order

`POST /v1/climate/orders/{order}`

Operation ID: `PostClimateOrdersOrder`

<p>Updates the specified order by setting the values of the parameters passed.</p>

## Definition

```yaml
{"summary": "Update an order", "description": "<p>Updates the specified order by setting the values of the parameters passed.</p>", "operationId": "PostClimateOrdersOrder", "parameters": [{"name": "order", "in": "path", "description": "Unique identifier of the order.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"beneficiary": {"description": "Publicly sharable reference for the end beneficiary of carbon removal. Assumed to be the Stripe account if not set.", "anyOf": [{"title": "beneficiary_params", "required": ["public_name"], "type": "object", "properties": {"public_name": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}}}, {"type": "string", "enum": [""]}]}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}}, "additionalProperties": false}, "encoding": {"beneficiary": {"style": "deepObject", "explode": true}, "expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/climate.order"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
