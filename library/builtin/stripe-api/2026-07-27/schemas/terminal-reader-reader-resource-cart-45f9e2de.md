---
title: terminal_reader_reader_resource_cart
page_id: schema-terminal-reader-reader-resource-cart-45f9e2de
path: schemas
description: Represents a cart to be displayed on the reader
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_reader_reader_resource_cart

Represents a cart to be displayed on the reader

```yaml
{"title": "TerminalReaderReaderResourceCart", "required": ["currency", "line_items", "total"], "type": "object", "properties": {"currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "line_items": {"type": "array", "description": "List of line items in the cart.", "items": {"$ref": "#/components/schemas/terminal_reader_reader_resource_line_item"}}, "tax": {"type": "integer", "description": "Tax amount for the entire cart. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).", "nullable": true}, "total": {"type": "integer", "description": "Total amount for the entire cart, including tax. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal)."}}, "description": "Represents a cart to be displayed on the reader", "x-expandableFields": ["line_items"]}
```
