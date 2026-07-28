---
title: Bulk remove invoice line items
page_id: operation-post-v1-invoices-invoice-remove-lines-7284c534
path: operations/untagged
description: <p>Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/invoices/{invoice}/remove_lines
operation_ids:
    - PostInvoicesInvoiceRemoveLines
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Bulk remove invoice line items

`POST /v1/invoices/{invoice}/remove_lines`

Operation ID: `PostInvoicesInvoiceRemoveLines`

<p>Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.</p>

## Definition

```yaml
{"summary": "Bulk remove invoice line items", "description": "<p>Removes multiple line items from an invoice. This is only possible when an invoice is still a draft.</p>", "operationId": "PostInvoicesInvoiceRemoveLines", "parameters": [{"name": "invoice", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"required": ["lines"], "type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "invoice_metadata": {"description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.", "anyOf": [{"type": "object", "additionalProperties": {"type": "string"}}, {"type": "string", "enum": [""]}]}, "lines": {"type": "array", "description": "The line items to remove.", "items": {"title": "lines_data_param", "required": ["behavior", "id"], "type": "object", "properties": {"behavior": {"type": "string", "enum": ["delete", "unassign"]}, "id": {"maxLength": 5000, "type": "string"}}}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "invoice_metadata": {"style": "deepObject", "explode": true}, "lines": {"style": "deepObject", "explode": true}}}}, "required": true}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/invoice"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
