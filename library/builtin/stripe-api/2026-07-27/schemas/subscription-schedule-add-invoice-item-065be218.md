---
title: subscription_schedule_add_invoice_item
page_id: schema-subscription-schedule-add-invoice-item-065be218
path: schemas
description: An Add Invoice Item describes the prices and quantities that will be added as pending invoice items when entering a phase.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedule_add_invoice_item

An Add Invoice Item describes the prices and quantities that will be added as pending invoice items when entering a phase.

```yaml
{"title": "SubscriptionScheduleAddInvoiceItem", "required": ["discounts", "period", "price"], "type": "object", "properties": {"discountable": {"type": "boolean", "description": "Controls whether discounts apply to this invoice item. Defaults to true if no value is provided.", "nullable": true}, "discounts": {"type": "array", "description": "The stackable discounts that will be applied to the item.", "items": {"$ref": "#/components/schemas/discounts_resource_stackable_discount_with_discount_end"}}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "period": {"$ref": "#/components/schemas/subscription_schedule_add_invoice_item_period"}, "price": {"description": "ID of the price used to generate the invoice item.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/price"}, {"$ref": "#/components/schemas/deleted_price"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/price"}, {"$ref": "#/components/schemas/deleted_price"}]}}, "quantity": {"type": "integer", "description": "The quantity of the invoice item.", "nullable": true}, "tax_rates": {"type": "array", "description": "The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.", "nullable": true, "items": {"$ref": "#/components/schemas/tax_rate"}}}, "description": "An Add Invoice Item describes the prices and quantities that will be added as pending invoice items when entering a phase.", "x-expandableFields": ["discounts", "period", "price", "tax_rates"]}
```
