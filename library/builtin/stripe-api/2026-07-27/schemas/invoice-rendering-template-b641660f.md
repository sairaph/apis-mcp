---
title: invoice_rendering_template
page_id: schema-invoice-rendering-template-b641660f
path: schemas
description: |-
    Invoice Rendering Templates are used to configure how invoices are rendered on surfaces like the PDF. Invoice Rendering Templates
    can be created from within the Dashboard, and they can be used over the API when creating invoices.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_rendering_template

Invoice Rendering Templates are used to configure how invoices are rendered on surfaces like the PDF. Invoice Rendering Templates
can be created from within the Dashboard, and they can be used over the API when creating invoices.

```yaml
{"title": "InvoiceRenderingTemplate", "required": ["created", "id", "livemode", "object", "status", "version"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "nickname": {"maxLength": 5000, "type": "string", "description": "A brief description of the template, hidden from customers", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["invoice_rendering_template"]}, "status": {"type": "string", "description": "The status of the template, one of `active` or `archived`.", "enum": ["active", "archived"]}, "version": {"type": "integer", "description": "Version of this template; version increases by one when an update on the template changes any field that controls invoice rendering"}}, "description": "Invoice Rendering Templates are used to configure how invoices are rendered on surfaces like the PDF. Invoice Rendering Templates\ncan be created from within the Dashboard, and they can be used over the API when creating invoices.", "x-expandableFields": [], "x-resourceId": "invoice_rendering_template"}
```
