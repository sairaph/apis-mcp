---
title: shipping_rate_currency_option
page_id: schema-shipping-rate-currency-option-d98423c0
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# shipping_rate_currency_option

```yaml
{"title": "ShippingRateCurrencyOption", "required": ["amount", "tax_behavior"], "type": "object", "properties": {"amount": {"type": "integer", "description": "A non-negative integer in cents representing how much to charge."}, "tax_behavior": {"type": "string", "description": "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.", "enum": ["exclusive", "inclusive", "unspecified"]}}, "description": "", "x-expandableFields": []}
```
