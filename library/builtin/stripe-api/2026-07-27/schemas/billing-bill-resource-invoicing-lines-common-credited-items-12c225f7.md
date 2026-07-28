---
title: billing_bill_resource_invoicing_lines_common_credited_items
page_id: schema-billing-bill-resource-invoicing-lines-common-credited-items-12c225f7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_lines_common_credited_items

```yaml
{"title": "BillingBillResourceInvoicingLinesCommonCreditedItems", "required": ["invoice", "invoice_line_items"], "type": "object", "properties": {"invoice": {"maxLength": 5000, "type": "string", "description": "Invoice containing the credited invoice line items"}, "invoice_line_items": {"type": "array", "description": "Credited invoice line items", "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
