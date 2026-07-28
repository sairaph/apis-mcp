---
title: source_order_item
page_id: schema-source-order-item-cd777ba7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_order_item

```yaml
{"title": "SourceOrderItem", "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount (price) for this order item.", "nullable": true}, "currency": {"maxLength": 5000, "type": "string", "description": "This currency of this order item. Required when `amount` is present.", "nullable": true}, "description": {"maxLength": 5000, "type": "string", "description": "Human-readable description for this order item.", "nullable": true}, "parent": {"maxLength": 5000, "type": "string", "description": "The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).", "nullable": true}, "quantity": {"type": "integer", "description": "The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered."}, "type": {"maxLength": 5000, "type": "string", "description": "The type of this order item. Must be `sku`, `tax`, or `shipping`.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
