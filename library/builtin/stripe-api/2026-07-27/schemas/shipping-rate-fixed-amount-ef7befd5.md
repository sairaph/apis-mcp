---
title: shipping_rate_fixed_amount
page_id: schema-shipping-rate-fixed-amount-ef7befd5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# shipping_rate_fixed_amount

```yaml
{"title": "ShippingRateFixedAmount", "required": ["amount", "currency"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A non-negative integer in cents representing how much to charge."}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).", "format": "currency"}, "currency_options": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/shipping_rate_currency_option"}, "description": "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies)."}}, "description": "", "x-expandableFields": ["currency_options"]}
```
