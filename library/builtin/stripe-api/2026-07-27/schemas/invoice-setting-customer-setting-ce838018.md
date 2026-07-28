---
title: invoice_setting_customer_setting
page_id: schema-invoice-setting-customer-setting-ce838018
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_setting_customer_setting

```yaml
{"title": "InvoiceSettingCustomerSetting", "type": "object", "properties": {"custom_fields": {"type": "array", "description": "Default custom fields to be displayed on invoices for this customer.", "nullable": true, "items": {"$ref": "#/components/schemas/invoice_setting_custom_field"}}, "default_payment_method": {"description": "ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_method"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_method"}]}}, "footer": {"maxLength": 5000, "type": "string", "description": "Default footer to be displayed on invoices for this customer.", "nullable": true}, "rendering_options": {"description": "Default options for invoice PDF rendering for this customer.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/invoice_setting_customer_rendering_options"}]}}, "description": "", "x-expandableFields": ["custom_fields", "default_payment_method", "rendering_options"]}
```
