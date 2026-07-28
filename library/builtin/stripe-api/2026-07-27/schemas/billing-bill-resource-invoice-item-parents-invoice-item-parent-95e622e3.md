---
title: billing_bill_resource_invoice_item_parents_invoice_item_parent
page_id: schema-billing-bill-resource-invoice-item-parents-invoice-item-parent-95e622e3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoice_item_parents_invoice_item_parent

```yaml
{"title": "BillingBillResourceInvoiceItemParentsInvoiceItemParent", "required": ["type"], "type": "object", "properties": {"subscription_details": {"description": "Details about the subscription that generated this invoice item", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_bill_resource_invoice_item_parents_invoice_item_subscription_parent"}]}, "type": {"type": "string", "description": "The type of parent that generated this invoice item", "enum": ["subscription_details"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": ["subscription_details"]}
```
