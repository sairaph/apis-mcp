---
title: tax_product_resource_line_item_tax_rate_details
page_id: schema-tax-product-resource-line-item-tax-rate-details-35646aac
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# tax_product_resource_line_item_tax_rate_details

```yaml
{"title": "TaxProductResourceLineItemTaxRateDetails", "required": ["display_name", "percentage_decimal", "tax_type"], "type": "object", "properties": {"display_name": {"maxLength": 5000, "type": "string", "description": "A localized display name for tax type, intended to be human-readable. For example, \"Local Sales and Use Tax\", \"Value-added tax (VAT)\", or \"Umsatzsteuer (USt.)\"."}, "percentage_decimal": {"maxLength": 5000, "type": "string", "description": "The tax rate percentage as a string. For example, 8.5% is represented as \"8.5\"."}, "tax_type": {"type": "string", "description": "The tax type, such as `vat` or `sales_tax`.", "enum": ["amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
