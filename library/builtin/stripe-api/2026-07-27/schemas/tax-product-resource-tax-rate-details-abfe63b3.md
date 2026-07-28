---
title: tax_product_resource_tax_rate_details
page_id: schema-tax-product-resource-tax-rate-details-abfe63b3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_rate_details

```yaml
{"title": "TaxProductResourceTaxRateDetails", "required": ["percentage_decimal"], "type": "object", "properties": {"country": {"maxLength": 5000, "type": "string", "description": "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).", "nullable": true}, "flat_amount": {"description": "The amount of the tax rate when the `rate_type` is `flat_amount`. Tax rates with `rate_type` `percentage` can vary based on the transaction, resulting in this field being `null`. This field exposes the amount and currency of the flat tax rate.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/tax_rate_flat_amount"}]}, "percentage_decimal": {"maxLength": 5000, "type": "string", "description": "The tax rate percentage as a string. For example, 8.5% is represented as `\"8.5\"`."}, "rate_type": {"type": "string", "description": "Indicates the type of tax rate applied to the taxable amount. This value can be `null` when no tax applies to the location. This field is only present for TaxRates created by Stripe Tax.", "nullable": true, "enum": ["flat_amount", "percentage"]}, "state": {"maxLength": 5000, "type": "string", "description": "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).", "nullable": true}, "tax_type": {"type": "string", "description": "The tax type, such as `vat` or `sales_tax`.", "nullable": true, "enum": ["amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["flat_amount"]}
```
