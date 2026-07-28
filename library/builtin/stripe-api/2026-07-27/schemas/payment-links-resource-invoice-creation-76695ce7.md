---
title: payment_links_resource_invoice_creation
page_id: schema-payment-links-resource-invoice-creation-76695ce7
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_invoice_creation

```yaml
{"title": "PaymentLinksResourceInvoiceCreation", "required": ["enabled"], "type": "object", "properties": {"enabled": {"type": "boolean", "description": "Enable creating an invoice on successful payment."}, "invoice_data": {"description": "Configuration for the invoice. Default invoice values will be used if unspecified.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/payment_links_resource_invoice_settings"}]}}, "description": "", "x-expandableFields": ["invoice_data"]}
```
