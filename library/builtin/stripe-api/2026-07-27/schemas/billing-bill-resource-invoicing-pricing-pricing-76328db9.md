---
title: billing_bill_resource_invoicing_pricing_pricing
page_id: schema-billing-bill-resource-invoicing-pricing-pricing-76328db9
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_pricing_pricing

```yaml
{"title": "BillingBillResourceInvoicingPricingPricing", "required": ["type"], "type": "object", "properties": {"price_details": {"$ref": "#/components/schemas/billing_bill_resource_invoicing_pricing_pricing_price_details"}, "type": {"type": "string", "description": "The type of the pricing details.", "enum": ["price_details"], "x-stripeBypassValidation": true}, "unit_amount_decimal": {"type": "string", "description": "The unit amount (in the `currency` specified) of the item which contains a decimal value with at most 12 decimal places.", "format": "decimal", "nullable": true}}, "description": "", "x-expandableFields": ["price_details"]}
```
