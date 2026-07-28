---
title: Remove a feature from a product
page_id: operation-delete-v1-products-product-features-id-3fe4e8f5
path: operations/untagged
description: <p>Deletes the feature attachment to a product</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/products/{product}/features/{id}
operation_ids:
    - DeleteProductsProductFeaturesId
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Remove a feature from a product

`DELETE /v1/products/{product}/features/{id}`

Operation ID: `DeleteProductsProductFeaturesId`

<p>Deletes the feature attachment to a product</p>

## Definition

```yaml
{"summary": "Remove a feature from a product", "description": "<p>Deletes the feature attachment to a product</p>", "operationId": "DeleteProductsProductFeaturesId", "parameters": [{"name": "id", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "product", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_product_feature"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
