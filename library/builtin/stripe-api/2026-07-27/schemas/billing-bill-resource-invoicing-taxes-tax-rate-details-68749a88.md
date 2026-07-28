---
title: billing_bill_resource_invoicing_taxes_tax_rate_details
page_id: schema-billing-bill-resource-invoicing-taxes-tax-rate-details-68749a88
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_taxes_tax_rate_details

```yaml
{"title": "BillingBillResourceInvoicingTaxesTaxRateDetails", "required": ["tax_rate"], "type": "object", "properties": {"tax_rate": {"description": "ID of the tax rate", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/tax_rate"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/tax_rate"}]}}}, "description": "", "x-expandableFields": ["tax_rate"]}
```
