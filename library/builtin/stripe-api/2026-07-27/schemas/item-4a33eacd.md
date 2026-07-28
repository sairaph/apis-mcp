---
title: item
page_id: schema-item-4a33eacd
path: schemas
description: A line item.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# item

A line item.

```yaml
{"title": "LineItem", "required": ["amount_discount", "amount_subtotal", "amount_tax", "amount_total", "currency", "id", "object"], "type": "object", "properties": {"adjustable_quantity": {"nullable": true, "anyOf": [{"$ref": "#/components/schemas/line_items_adjustable_quantity"}]}, "amount_discount": {"type": "integer", "description": "Total discount amount applied. If no discounts were applied, defaults to 0."}, "amount_subtotal": {"type": "integer", "description": "Total before any discounts or taxes are applied."}, "amount_tax": {"type": "integer", "description": "Total tax amount applied. If no tax was applied, defaults to 0."}, "amount_total": {"type": "integer", "description": "Total after discounts and taxes."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users. Defaults to product name.", "nullable": true}, "discounts": {"type": "array", "description": "The discounts applied to the line item.", "items": {"$ref": "#/components/schemas/line_items_discount_amount"}}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["item"]}, "price": {"description": "The price used to generate the line item.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/price"}]}, "quantity": {"type": "integer", "description": "The quantity of products being purchased.", "nullable": true}, "taxes": {"type": "array", "description": "The taxes applied to the line item.", "items": {"$ref": "#/components/schemas/line_items_tax_amount"}}}, "description": "A line item.", "x-expandableFields": ["adjustable_quantity", "discounts", "price", "taxes"], "x-resourceId": "item"}
```
