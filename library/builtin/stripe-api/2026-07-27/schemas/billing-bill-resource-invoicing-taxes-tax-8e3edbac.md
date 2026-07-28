---
title: billing_bill_resource_invoicing_taxes_tax
page_id: schema-billing-bill-resource-invoicing-taxes-tax-8e3edbac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_taxes_tax

```yaml
{"title": "BillingBillResourceInvoicingTaxesTax", "required": ["amount", "tax_behavior", "taxability_reason", "type"], "type": "object", "properties": {"amount": {"type": "integer", "description": "The amount of the tax, in cents (or local equivalent)."}, "tax_behavior": {"type": "string", "description": "Whether this tax is inclusive or exclusive.", "enum": ["exclusive", "inclusive"]}, "tax_rate_details": {"description": "Additional details about the tax rate. Only present when `type` is `tax_rate_details`.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_taxes_tax_rate_details"}]}, "taxability_reason": {"type": "string", "description": "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.", "enum": ["customer_exempt", "not_available", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated"], "x-stripeBypassValidation": true}, "taxable_amount": {"type": "integer", "description": "The amount on which tax is calculated, in cents (or local equivalent).", "nullable": true}, "type": {"type": "string", "description": "The type of tax information.", "enum": ["tax_rate_details"]}}, "description": "", "x-expandableFields": ["tax_rate_details"]}
```
