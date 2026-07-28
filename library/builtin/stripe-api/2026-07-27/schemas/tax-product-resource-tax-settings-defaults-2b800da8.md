---
title: tax_product_resource_tax_settings_defaults
page_id: schema-tax-product-resource-tax-settings-defaults-2b800da8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_settings_defaults

```yaml
{"title": "TaxProductResourceTaxSettingsDefaults", "required": ["provider"], "type": "object", "properties": {"provider": {"type": "string", "description": "The tax calculation provider this account uses. Defaults to `stripe` when not using a [third-party provider](/tax/third-party-apps).", "enum": ["anrok", "avalara", "sphere", "stripe"]}, "tax_behavior": {"type": "string", "description": "Default [tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#tax-behavior) used to specify whether the price is considered inclusive of taxes or exclusive of taxes. If the item's price has a tax behavior set, it will take precedence over the default tax behavior.", "nullable": true, "enum": ["exclusive", "inclusive", "inferred_by_currency"]}, "tax_code": {"maxLength": 5000, "type": "string", "description": "Default [tax code](https://stripe.com/docs/tax/tax-categories) used to classify your products and prices.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
