---
title: subscription_schedule_configuration_item
page_id: schema-subscription-schedule-configuration-item-7ddee2b1
path: schemas
description: A phase item describes the price and quantity of a phase.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedule_configuration_item

A phase item describes the price and quantity of a phase.

```yaml
{"title": "SubscriptionScheduleConfigurationItem", "required": ["discounts", "price"], "type": "object", "properties": {"billing_thresholds": {"description": "Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscription_item_billing_thresholds"}]}, "discounts": {"type": "array", "description": "The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.", "items": {"$ref": "#/components/schemas/stackable_discount_with_discount_settings"}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.", "nullable": true}, "price": {"description": "ID of the price to which the customer should be subscribed.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/price"}, {"$ref": "#/components/schemas/deleted_price"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/price"}, {"$ref": "#/components/schemas/deleted_price"}]}}, "quantity": {"type": "integer", "description": "Quantity of the plan to which the customer should be subscribed."}, "tax_rates": {"type": "array", "description": "The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.", "nullable": true, "items": {"$ref": "#/components/schemas/tax_rate"}}}, "description": "A phase item describes the price and quantity of a phase.", "x-expandableFields": ["billing_thresholds", "discounts", "price", "tax_rates"]}
```
