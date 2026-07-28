---
title: subscription_item
page_id: schema-subscription-item-545fd055
path: schemas
description: |-
    Subscription items allow you to create customer subscriptions with more than
    one plan, making it easy to represent complex billing relationships.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_item

Subscription items allow you to create customer subscriptions with more than
one plan, making it easy to represent complex billing relationships.

```yaml
{"title": "SubscriptionItem", "required": ["created", "current_period_end", "current_period_start", "discounts", "id", "metadata", "object", "price", "subscription"], "type": "object", "properties": {"billed_until": {"type": "integer", "description": "The time period the subscription item has been billed for.", "format": "unix-time"}, "billing_thresholds": {"description": "Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscription_item_billing_thresholds"}]}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch."}, "current_period_end": {"type": "integer", "description": "The end time of this subscription item's current billing period.", "format": "unix-time"}, "current_period_start": {"type": "integer", "description": "The start time of this subscription item's current billing period.", "format": "unix-time"}, "discounts": {"type": "array", "description": "The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.", "items": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}]}}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["subscription_item"]}, "price": {"$ref": "#/components/schemas/price"}, "quantity": {"type": "integer", "description": "The [quantity](https://docs.stripe.com/subscriptions/quantities) of the plan to which the customer should be subscribed."}, "subscription": {"maxLength": 5000, "type": "string", "description": "The `subscription` this `subscription_item` belongs to."}, "tax_rates": {"type": "array", "description": "The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.", "nullable": true, "items": {"$ref": "#/components/schemas/tax_rate"}}}, "description": "Subscription items allow you to create customer subscriptions with more than\none plan, making it easy to represent complex billing relationships.", "x-expandableFields": ["billing_thresholds", "discounts", "price", "tax_rates"], "x-resourceId": "subscription_item"}
```
