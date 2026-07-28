---
title: billing_bill_resource_invoicing_lines_parents_invoice_line_item_invoice_item_parent
page_id: schema-billing-bill-resource-invoicing-lines-parents-invoice-line-item-invoice-379df44c
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_lines_parents_invoice_line_item_invoice_item_parent

```yaml
{"title": "BillingBillResourceInvoicingLinesParentsInvoiceLineItemInvoiceItemParent", "required": ["invoice_item", "proration"], "type": "object", "properties": {"invoice_item": {"maxLength": 5000, "type": "string", "description": "The invoice item that generated this line item"}, "proration": {"type": "boolean", "description": "Whether this is a proration"}, "proration_details": {"description": "Additional details for proration line items", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoicing_lines_common_proration_details"}]}, "subscription": {"maxLength": 5000, "type": "string", "description": "The subscription that the invoice item belongs to", "nullable": true}}, "description": "", "x-expandableFields": ["proration_details"]}
```
