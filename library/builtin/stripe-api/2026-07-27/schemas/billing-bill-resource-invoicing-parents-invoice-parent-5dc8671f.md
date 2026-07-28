---
title: billing_bill_resource_invoicing_parents_invoice_parent
page_id: schema-billing-bill-resource-invoicing-parents-invoice-parent-5dc8671f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_parents_invoice_parent

```yaml
{"title": "BillingBillResourceInvoicingParentsInvoiceParent", "required": ["type"], "type": "object", "properties": {"quote_details": {"description": "Details about the quote that generated this invoice", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_parents_invoice_quote_parent"}]}, "subscription_details": {"description": "Details about the subscription that generated this invoice", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_parents_invoice_subscription_parent"}]}, "type": {"type": "string", "description": "The type of parent that generated this invoice", "enum": ["quote_details", "subscription_details"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["quote_details", "subscription_details"]}
```
