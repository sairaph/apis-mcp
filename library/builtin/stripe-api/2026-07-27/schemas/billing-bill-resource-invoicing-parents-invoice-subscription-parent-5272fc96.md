---
title: billing_bill_resource_invoicing_parents_invoice_subscription_parent
page_id: schema-billing-bill-resource-invoicing-parents-invoice-subscription-parent-5272fc96
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_bill_resource_invoicing_parents_invoice_subscription_parent

```yaml
{"title": "BillingBillResourceInvoicingParentsInvoiceSubscriptionParent", "required": ["subscription"], "type": "object", "properties": {"metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) defined as subscription metadata when an invoice is created. Becomes an immutable snapshot of the subscription metadata at the time of invoice finalization.\n *Note: This attribute is populated only for invoices created on or after June 29, 2023.*", "nullable": true}, "subscription": {"description": "The subscription that generated this invoice", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/subscription"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/subscription"}]}}, "subscription_proration_date": {"type": "integer", "description": "Only set for upcoming invoices that preview prorations. The time used to calculate prorations.", "format": "unix-time"}}, "description": "", "x-expandableFields": ["subscription"]}
```
