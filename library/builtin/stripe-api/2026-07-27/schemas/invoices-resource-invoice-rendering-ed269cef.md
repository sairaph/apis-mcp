---
title: invoices_resource_invoice_rendering
page_id: schema-invoices-resource-invoice-rendering-ed269cef
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_resource_invoice_rendering

```yaml
{"title": "InvoicesResourceInvoiceRendering", "type": "object", "properties": {"amount_tax_display": {"maxLength": 5000, "type": "string", "description": "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.", "nullable": true}, "pdf": {"description": "Invoice pdf rendering options", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/invoice_rendering_pdf"}]}, "template": {"maxLength": 5000, "type": "string", "description": "ID of the rendering template that the invoice is formatted by.", "nullable": true}, "template_version": {"type": "integer", "description": "Version of the rendering template that the invoice is using.", "nullable": true}}, "description": "", "x-expandableFields": ["pdf"]}
```
