---
title: tax_product_resource_line_item_tax_breakdown
page_id: schema-tax-product-resource-line-item-tax-breakdown-2eed6a20
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_line_item_tax_breakdown

```yaml
{"title": "TaxProductResourceLineItemTaxBreakdown", "required": ["amount", "jurisdiction", "sourcing", "taxability_reason", "taxable_amount"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount of tax, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}, "jurisdiction": {"$ref": "#/components/schemas/tax_product_resource_jurisdiction"}, "sourcing": {"type": "string", "description": "Indicates whether the jurisdiction was determined by the origin (merchant's address) or destination (customer's address).", "enum": ["destination", "origin"], "x-stripeBypassValidation": true}, "tax_rate_details": {"description": "Details regarding the rate for this tax. This field will be `null` when the tax is not imposed, for example if the product is exempt from tax.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/tax_product_resource_line_item_tax_rate_details"}]}, "taxability_reason": {"type": "string", "description": "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.", "enum": ["customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated"]}, "taxable_amount": {"type": "integer", "description": "The amount on which tax is calculated, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units)."}}, "description": "", "x-expandableFields": ["jurisdiction", "tax_rate_details"]}
```
