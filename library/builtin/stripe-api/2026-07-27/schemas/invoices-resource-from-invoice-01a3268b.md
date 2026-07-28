---
title: invoices_resource_from_invoice
page_id: schema-invoices-resource-from-invoice-01a3268b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_from_invoice

```yaml
{"title": "InvoicesResourceFromInvoice", "required": ["action", "invoice"], "type": "object", "properties": {"action": {"maxLength": 5000, "type": "string", "description": "The relation between this invoice and the cloned invoice"}, "invoice": {"description": "The invoice that was cloned.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/invoice"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/invoice"}]}}}, "description": "", "x-expandableFields": ["invoice"]}
```
