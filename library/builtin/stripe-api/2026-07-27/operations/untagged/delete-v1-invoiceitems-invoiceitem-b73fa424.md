---
title: Delete an invoice item
page_id: operation-delete-v1-invoiceitems-invoiceitem-6f93aedb
path: operations/untagged
description: <p>Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they’re not attached to invoices, or if it’s attached to a draft invoice.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/invoiceitems/{invoiceitem}
operation_ids:
    - DeleteInvoiceitemsInvoiceitem
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete an invoice item

`DELETE /v1/invoiceitems/{invoiceitem}`

Operation ID: `DeleteInvoiceitemsInvoiceitem`

<p>Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they’re not attached to invoices, or if it’s attached to a draft invoice.</p>

## Definition

```yaml
{"summary": "Delete an invoice item", "description": "<p>Deletes an invoice item, removing it from an invoice. Deleting invoice items is only possible when they’re not attached to invoices, or if it’s attached to a draft invoice.</p>", "operationId": "DeleteInvoiceitemsInvoiceitem", "parameters": [{"name": "invoiceitem", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_invoiceitem"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
