---
title: tax_product_resource_tax_breakdown
page_id: schema-tax-product-resource-tax-breakdown-9ee75d99
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_tax_breakdown

```yaml
{"title": "TaxProductResourceTaxBreakdown", "required": ["amount", "inclusive", "tax_rate_details", "taxability_reason", "taxable_amount"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount of tax, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}, "inclusive": {"type": "boolean", "description": "Specifies whether the tax amount is included in the line item amount."}, "tax_rate_details": {"$ref": "#/components/schemas/tax_product_resource_tax_rate_details"}, "taxability_reason": {"type": "string", "description": "The reasoning behind this tax, for example, if the product is tax exempt. We might extend the possible values for this field to support new tax rules.", "enum": ["customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated"]}, "taxable_amount": {"type": "integer", "description": "The amount on which tax is calculated, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}}, "description": "", "x-expandableFields": ["tax_rate_details"]}
```
