---
title: Attach a feature to a product
page_id: operation-post-v1-products-product-features-937c840f
path: operations/untagged
description: <p>Creates a product_feature, which represents a feature attachment to a product</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/products/{product}/features
operation_ids:
    - PostProductsProductFeatures
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Attach a feature to a product

`POST /v1/products/{product}/features`

Operation ID: `PostProductsProductFeatures`

<p>Creates a product_feature, which represents a feature attachment to a product</p>

## Definition

```yaml
{"summary": "Attach a feature to a product", "description": "<p>Creates a product_feature, which represents a feature attachment to a product</p>", "operationId": "PostProductsProductFeatures", "parameters": [{"name": "product", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["entitlement_feature"], "type": "object", "properties": {"entitlement_feature": {"maxLength": 5000, "type": "string", "description": "The ID of the [Feature](https://docs.stripe.com/api/entitlements/feature) object attached to this product."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/product_feature"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
