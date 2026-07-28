---
title: Retrieve a coupon
page_id: operation-get-v1-coupons-coupon-8a5d366b
path: operations/untagged
description: <p>Retrieves the coupon with the given ID.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/coupons/{coupon}
operation_ids:
    - GetCouponsCoupon
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a coupon

`GET /v1/coupons/{coupon}`

Operation ID: `GetCouponsCoupon`

<p>Retrieves the coupon with the given ID.</p>

## Definition

```yaml
{"summary": "Retrieve a coupon", "description": "<p>Retrieves the coupon with the given ID.</p>", "operationId": "GetCouponsCoupon", "parameters": [{"name": "coupon", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}, {"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/coupon"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
