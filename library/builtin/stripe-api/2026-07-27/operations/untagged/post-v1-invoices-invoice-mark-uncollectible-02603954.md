---
title: Mark an invoice as uncollectible
page_id: operation-post-v1-invoices-invoice-mark-uncollectible-7d937b10
path: operations/untagged
description: <p>Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/invoices/{invoice}/mark_uncollectible
operation_ids:
    - PostInvoicesInvoiceMarkUncollectible
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Mark an invoice as uncollectible

`POST /v1/invoices/{invoice}/mark_uncollectible`

Operation ID: `PostInvoicesInvoiceMarkUncollectible`

<p>Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.</p>

## Definition

```yaml
{"summary": "Mark an invoice as uncollectible", "description": "<p>Marking an invoice as uncollectible is useful for keeping track of bad debts that can be written off for accounting purposes.</p>", "operationId": "PostInvoicesInvoiceMarkUncollectible", "parameters": [{"name": "invoice", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/invoice"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
