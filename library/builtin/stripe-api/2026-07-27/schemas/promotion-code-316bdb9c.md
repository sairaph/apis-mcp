---
title: promotion_code
page_id: schema-promotion-code-316bdb9c
path: schemas
description: |-
    A Promotion Code represents a customer-redeemable code for an underlying promotion.
    You can create multiple codes for a single promotion.

    If you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.
    Customers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# promotion_code

A Promotion Code represents a customer-redeemable code for an underlying promotion.
You can create multiple codes for a single promotion.

If you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.
Customers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.

```yaml
{"title": "PromotionCode", "required": ["active", "code", "created", "id", "livemode", "object", "promotion", "restrictions", "times_redeemed"], "type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the promotion code is currently active. A promotion code is only active if the coupon is also valid."}, "code": {"maxLength": 5000, "type": "string", "description": "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for each customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-)."}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "customer": {"description": "The customer who can use this promotion code.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}, {"$ref": "#/components/schemas/deleted_customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}, {"$ref": "#/components/schemas/deleted_customer"}]}}, "customer_account": {"maxLength": 5000, "type": "string", "description": "The account representing the customer who can use this promotion code.", "nullable": true}, "expires_at": {"type": "integer", "description": "Date at which the promotion code can no longer be redeemed.", "format": "unix-time", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "max_redemptions": {"type": "integer", "description": "Maximum number of times this promotion code can be redeemed.", "nullable": true}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["promotion_code"]}, "promotion": {"$ref": "#/components/schemas/promotion_codes_resource_promotion"}, "restrictions": {"$ref": "#/components/schemas/promotion_codes_resource_restrictions"}, "times_redeemed": {"type": "integer", "description": "Number of times this promotion code has been used."}}, "description": "A Promotion Code represents a customer-redeemable code for an underlying promotion.\nYou can create multiple codes for a single promotion.\n\nIf you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.\nCustomers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.", "x-expandableFields": ["customer", "promotion", "restrictions"], "x-resourceId": "promotion_code"}
```
