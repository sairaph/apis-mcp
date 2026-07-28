---
title: payment_links_resource_optional_item_adjustable_quantity
page_id: schema-payment-links-resource-optional-item-adjustable-quantity-50537a22
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_optional_item_adjustable_quantity

```yaml
{"title": "PaymentLinksResourceOptionalItemAdjustableQuantity", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Set to true if the quantity can be adjusted to any non-negative integer."}, "maximum": {"type": "integer", "description": "The maximum quantity of this item the customer can purchase. By default this value is 99.", "nullable": true}, "minimum": {"type": "integer", "description": "The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
