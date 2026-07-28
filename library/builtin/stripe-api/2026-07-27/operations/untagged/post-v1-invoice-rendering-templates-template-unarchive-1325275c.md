---
title: Unarchive an invoice rendering template
page_id: operation-post-v1-invoice-rendering-templates-template-unarchive-22e5a97e
path: operations/untagged
description: <p>Unarchive an invoice rendering template so it can be used on new Stripe objects again.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/invoice_rendering_templates/{template}/unarchive
operation_ids:
    - PostInvoiceRenderingTemplatesTemplateUnarchive
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Unarchive an invoice rendering template

`POST /v1/invoice_rendering_templates/{template}/unarchive`

Operation ID: `PostInvoiceRenderingTemplatesTemplateUnarchive`

<p>Unarchive an invoice rendering template so it can be used on new Stripe objects again.</p>

## Definition

```yaml
{"summary": "Unarchive an invoice rendering template", "description": "<p>Unarchive an invoice rendering template so it can be used on new Stripe objects again.</p>", "operationId": "PostInvoiceRenderingTemplatesTemplateUnarchive", "parameters": [{"name": "template", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/invoice_rendering_template"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
