---
title: tax_product_resource_tax_transaction_shipping_cost
page_id: schema-tax-product-resource-tax-transaction-shipping-cost-63597eaa
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_transaction_shipping_cost

```yaml
{"title": "TaxProductResourceTaxTransactionShippingCost", "required": ["amount", "amount_tax", "tax_behavior", "tax_code"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The shipping amount in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units). If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes were calculated on top of this amount."}, "amount_tax": {"type": "integer", "description": "The amount of tax calculated for shipping, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}, "shipping_rate": {"maxLength": 5000, "type": "string", "description": "The ID of an existing [ShippingRate](https://docs.stripe.com/api/shipping_rates/object)."}, "tax_behavior": {"type": "string", "description": "Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes.", "enum": ["exclusive", "inclusive"]}, "tax_code": {"maxLength": 5000, "type": "string", "description": "The [tax code](https://docs.stripe.com/tax/tax-categories) ID used for shipping."}}, "description": "", "x-expandableFields": []}
```
