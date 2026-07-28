---
title: Retrieve a product
page_id: operation-get-v1-climate-products-product-3beead73
path: operations/untagged
description: <p>Retrieves the details of a Climate product with the given ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/climate/products/{product}
operation_ids:
    - GetClimateProductsProduct
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a product

`GET /v1/climate/products/{product}`

Operation ID: `GetClimateProductsProduct`

<p>Retrieves the details of a Climate product with the given ID.</p>

## Definition

```yaml
{"summary": "Retrieve a product", "description": "<p>Retrieves the details of a Climate product with the given ID.</p>", "operationId": "GetClimateProductsProduct", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "product", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/climate.product"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
