---
title: line_items_tax_amount
page_id: schema-line-items-tax-amount-a23c6725
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# line_items_tax_amount

```yaml
{"title": "LineItemsTaxAmount", "required": ["amount", "rate"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount of tax applied for this rate."}, "rate": {"$ref": "#/components/schemas/tax_rate"}, "taxability_reason": {"type": "string", "description": "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.", "nullable": true, "enum": ["customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated"], "x-stripeBypassValidation": true}, "taxable_amount": {"type": "integer", "description": "The amount on which tax is calculated, in cents (or local equivalent).", "nullable": true}}, "description": "", "x-expandableFields": ["rate"]}
```
