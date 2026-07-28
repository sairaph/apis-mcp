---
title: source_order
page_id: schema-source-order-cf3f7cfa
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# source_order

```yaml
{"title": "SourceOrder", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the order."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "email": {"maxLength": 5000, "type": "string", "description": "The email address of the customer placing the order."}, "items": {"type": "array", "description": "List of items constituting the order.", "nullable": true, "items": {"$ref": "#/components/schemas/source_order_item"}}, "shipping": {"$ref": "#/components/schemas/shipping"}}, "description": "", "x-expandableFields": ["items", "shipping"]}
```
