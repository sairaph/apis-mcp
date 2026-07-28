---
title: credited_items_invoice_line_items
page_id: schema-credited-items-invoice-line-items-7b5359a3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# credited_items_invoice_line_items

```yaml
{"title": "CreditedItemsInvoiceLineItems", "required": ["invoice", "invoice_line_items"], "type": "object", "properties": {"invoice": {"maxLength": 5000, "type": "string", "description": "The invoice id for the debited line item(s)."}, "invoice_line_items": {"type": "array", "description": "IDs of the debited invoice line item(s) on the invoice that correspond to the credit proration.", "items": {"maxLength": 5000, "type": "string"}}}, "description": "", "x-expandableFields": []}
```
