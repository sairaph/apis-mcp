---
title: shipping_rate
page_id: schema-shipping-rate-b6b4c1b1
path: schemas
description: |-
    Shipping rates describe the price of shipping presented to your customers and
    applied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# shipping_rate

Shipping rates describe the price of shipping presented to your customers and
applied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).

```yaml
{"title": "ShippingRate", "required": ["active", "created", "id", "livemode", "metadata", "object", "type"], "type": "object", "properties": {"active": {"type": "boolean", "description": "Whether the shipping rate can be used for new purchases. Defaults to `true`."}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "delivery_estimate": {"description": "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/shipping_rate_delivery_estimate"}]}, "display_name": {"maxLength": 5000, "type": "string", "description": "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.", "nullable": true}, "fixed_amount": {"$ref": "#/components/schemas/shipping_rate_fixed_amount"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["shipping_rate"]}, "tax_behavior": {"type": "string", "description": "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.", "nullable": true, "enum": ["exclusive", "inclusive", "unspecified"]}, "tax_code": {"description": "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/tax_code"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/tax_code"}]}}, "type": {"type": "string", "description": "The type of calculation to use on the shipping rate.", "enum": ["fixed_amount"]}}, "description": "Shipping rates describe the price of shipping presented to your customers and\napplied to a purchase. For more information, see [Charge for shipping](https://docs.stripe.com/payments/during-payment/charge-shipping).", "x-expandableFields": ["delivery_estimate", "fixed_amount", "tax_code"], "x-resourceId": "shipping_rate"}
```
