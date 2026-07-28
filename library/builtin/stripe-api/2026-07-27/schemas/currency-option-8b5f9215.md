---
title: currency_option
page_id: schema-currency-option-8b5f9215
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# currency_option

```yaml
{"title": "CurrencyOption", "type": "object", "properties": {"custom_unit_amount": {"description": "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/custom_unit_amount"}]}, "tax_behavior": {"type": "string", "description": "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.", "nullable": true, "enum": ["exclusive", "inclusive", "unspecified"]}, "tiers": {"type": "array", "description": "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.", "items": {"$ref": "#/components/schemas/price_tier"}}, "unit_amount": {"type": "integer", "description": "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.", "nullable": true}, "unit_amount_decimal": {"type": "string", "description": "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.", "format": "decimal", "nullable": true}}, "description": "", "x-expandableFields": ["custom_unit_amount", "tiers"]}
```
