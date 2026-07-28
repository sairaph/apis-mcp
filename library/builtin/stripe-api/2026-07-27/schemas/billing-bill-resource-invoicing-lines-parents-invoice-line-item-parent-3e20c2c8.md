---
title: billing_bill_resource_invoicing_lines_parents_invoice_line_item_parent
page_id: schema-billing-bill-resource-invoicing-lines-parents-invoice-line-item-parent-3e20c2c8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_lines_parents_invoice_line_item_parent

```yaml
{"title": "BillingBillResourceInvoicingLinesParentsInvoiceLineItemParent", "required": ["type"], "type": "object", "properties": {"invoice_item_details": {"description": "Details about the invoice item that generated this line item", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_lines_parents_invoice_line_item_invoice_item_parent"}]}, "subscription_item_details": {"description": "Details about the subscription item that generated this line item", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_lines_parents_invoice_line_item_subscription_item_parent"}]}, "type": {"type": "string", "description": "The type of parent that generated this line item", "enum": ["invoice_item_details", "subscription_item_details"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["invoice_item_details", "subscription_item_details"]}
```
