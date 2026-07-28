---
title: payment_links_resource_invoice_settings
page_id: schema-payment-links-resource-invoice-settings-13a39a27
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_links_resource_invoice_settings

```yaml
{"title": "PaymentLinksResourceInvoiceSettings", "type": "object", "properties": {"account_tax_ids": {"type": "array", "description": "The account tax IDs associated with the invoice.", "nullable": true, "items": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/tax_id"}, {"$ref": "#/components/schemas/deleted_tax_id"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/tax_id"}, {"$ref": "#/components/schemas/deleted_tax_id"}]}}}, "custom_fields": {"type": "array", "description": "A list of up to 4 custom fields to be displayed on the invoice.", "nullable": true, "items": {"$ref": "#/components/schemas/invoice_setting_custom_field"}}, "description": {"maxLength": 5000, "type": "string", "description": "An arbitrary string attached to the object. Often useful for displaying to users.", "nullable": true}, "footer": {"maxLength": 5000, "type": "string", "description": "Footer to be displayed on the invoice.", "nullable": true}, "issuer": {"description": "The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/connect_account_reference"}]}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "rendering_options": {"description": "Options for invoice PDF rendering.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/invoice_setting_checkout_rendering_options"}]}}, "description": "", "x-expandableFields": ["account_tax_ids", "custom_fields", "issuer", "rendering_options"]}
```
