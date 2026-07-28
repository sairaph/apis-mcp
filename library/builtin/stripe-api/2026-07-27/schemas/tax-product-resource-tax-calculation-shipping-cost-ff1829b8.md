---
title: tax_product_resource_tax_calculation_shipping_cost
page_id: schema-tax-product-resource-tax-calculation-shipping-cost-ff1829b8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_calculation_shipping_cost

```yaml
{"title": "TaxProductResourceTaxCalculationShippingCost", "required": ["amount", "amount_tax", "tax_behavior", "tax_code"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The shipping amount in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units). If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes were calculated on top of this amount."}, "amount_tax": {"type": "integer", "description": "The amount of tax calculated for shipping, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}, "shipping_rate": {"maxLength": 5000, "type": "string", "description": "The ID of an existing [ShippingRate](https://docs.stripe.com/api/shipping_rates/object)."}, "tax_behavior": {"type": "string", "description": "Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes.", "enum": ["exclusive", "inclusive"]}, "tax_breakdown": {"type": "array", "description": "Detailed account of taxes relevant to shipping cost.", "items": {"$ref": "#/components/schemas/tax_product_resource_line_item_tax_breakdown"}}, "tax_code": {"maxLength": 5000, "type": "string", "description": "The [tax code](https://docs.stripe.com/tax/tax-categories) ID used for shipping."}}, "description": "", "x-expandableFields": ["tax_breakdown"]}
```
