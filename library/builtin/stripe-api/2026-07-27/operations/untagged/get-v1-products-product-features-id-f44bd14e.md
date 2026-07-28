---
title: Retrieve a product_feature
page_id: operation-get-v1-products-product-features-id-e248f974
path: operations/untagged
description: <p>Retrieves a product_feature, which represents a feature attachment to a product</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/products/{product}/features/{id}
operation_ids:
    - GetProductsProductFeaturesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a product_feature

`GET /v1/products/{product}/features/{id}`

Operation ID: `GetProductsProductFeaturesId`

<p>Retrieves a product_feature, which represents a feature attachment to a product</p>

## Definition

```yaml
{"summary": "Retrieve a product_feature", "description": "<p>Retrieves a product_feature, which represents a feature attachment to a product</p>", "operationId": "GetProductsProductFeaturesId", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "id", "in": "path", "description": "The ID of the product_feature.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "product", "in": "path", "description": "The ID of the product.", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/product_feature"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
