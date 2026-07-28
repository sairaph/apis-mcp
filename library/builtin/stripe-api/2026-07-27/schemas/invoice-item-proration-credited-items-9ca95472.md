---
title: invoice_item_proration_credited_items
page_id: schema-invoice-item-proration-credited-items-9ca95472
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_item_proration_credited_items

```yaml
{"title": "InvoiceItemProrationCreditedItems", "required": ["type"], "type": "object", "properties": {"invoice_item": {"maxLength": 5000, "type": "string", "description": "When `type` is `invoice_item`, the invoice item id for the debited invoice item corresponding to this credit proration."}, "invoice_line_item_details": {"$ref": "#/components/schemas/credited_items_invoice_line_items"}, "type": {"type": "string", "description": "Whether the credit references a pending invoice item or one or more invoice line items on an invoice.", "enum": ["invoice_item", "invoice_line_items"]}}, "description": "", "x-expandableFields": ["invoice_line_item_details"]}
```
