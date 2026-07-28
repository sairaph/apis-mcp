---
title: Create a promotion code
page_id: operation-post-v1-promotion-codes-ade27888
path: operations/untagged
description: <p>A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/promotion_codes
operation_ids:
    - PostPromotionCodes
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Create a promotion code

`POST /v1/promotion_codes`

Operation ID: `PostPromotionCodes`

<p>A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.</p>

## Definition

```yaml
{"summary": "Create a promotion code", "description": "<p>A promotion code points to an underlying promotion. You can optionally restrict the code to a specific customer, redemption limit, and expiration date.</p>", "operationId": "PostPromotionCodes", "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["promotion"], "type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the promotion code is currently active."}, "code": {"maxLength": 500, "type": "string", "description": "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for a specific customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).\n\nIf left blank, we will generate one automatically."}, "customer": {"maxLength": 5000, "type": "string", "description": "The customer who can use this promotion code. If not set, all customers can use the promotion code."}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The account representing the customer who can use this promotion code. If not set, all customers can use the promotion code."}, "expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "expires_at": {"type": "integer", "description": "The timestamp at which this promotion code will expire. If the coupon has specified a `redeems_by`, then this value cannot be after the coupon's `redeems_by`.", "format": "unix-time"}, "max_redemptions": {"type": "integer", "description": "A positive integer specifying the number of times the promotion code can be redeemed. If the coupon has specified a `max_redemptions`, then this value cannot be greater than the coupon's `max_redemptions`."}, "metadata": {"type": "object", "additionalProperties": {"type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`."}, "promotion": {"title": "promotion", "required": ["type"], "type": "object", "properties": {"coupon": {"maxLength": 5000, "type": "string"}, "type": {"type": "string", "enum": ["coupon"]}}, "description": "The promotion referenced by this promotion code."}, "restrictions": {"title": "restrictions_params", "type": "object", "properties": {"currency_options": {"type": "object", "additionalProperties": {"title": "currency_option", "type": "object", "properties": {"minimum_amount": {"type": "integer"}}}}, "first_time_transaction": {"type": "boolean"}, "minimum_amount": {"type": "integer"}, "minimum_amount_currency": {"type": "string", "format": "currency"}}, "description": "Settings that restrict the redemption of the promotion code."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "metadata": {"style": "deepObject", "explode": true}, "promotion": {"style": "deepObject", "explode": true}, "restrictions": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/promotion_code"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
