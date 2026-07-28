---
title: Update a promotion code
page_id: operation-post-v1-promotion-codes-promotion-code-2ebc9aca
path: operations/untagged
description: <p>Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/promotion_codes/{promotion_code}
operation_ids:
    - PostPromotionCodesPromotionCode
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Update a promotion code

`POST /v1/promotion_codes/{promotion_code}`

Operation ID: `PostPromotionCodesPromotionCode`

<p>Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.</p>

## Definition

```yaml
{"summary": "Update a promotion code", "description": "<p>Updates the specified promotion code by setting the values of the parameters passed. Most fields are, by design, not editable.</p>", "operationId": "PostPromotionCodesPromotionCode", "parameters": [{"name": "promotion_code", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the promotion code is currently active. A promotion code can only be reactivated when the coupon is still valid and the promotion code is otherwise redeemable."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "restrictions": {"title": "restrictions_params", "type": "object", "properties": {"currency_options": {"type": "object", "additionalProperties": {"title": "currency_option", "type": "object", "properties": {"minimum_amount": {"type": "integer"}}}}}, "description": "Settings that restrict the redemption of the promotion code."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "restrictions": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/promotion_code"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
