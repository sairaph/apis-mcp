---
title: invoice_setting_customer_rendering_options
page_id: schema-invoice-setting-customer-rendering-options-f37bbd29
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_setting_customer_rendering_options

```yaml
{"title": "InvoiceSettingCustomerRenderingOptions", "type": "object", "properties": {"amount_tax_display": {"maxLength": 5000, "type": "string", "description": "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.", "nullable": true}, "template": {"maxLength": 5000, "type": "string", "description": "ID of the invoice rendering template to be used for this customer's invoices. If set, the template will be used on all invoices for this customer unless a template is set directly on the invoice.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
