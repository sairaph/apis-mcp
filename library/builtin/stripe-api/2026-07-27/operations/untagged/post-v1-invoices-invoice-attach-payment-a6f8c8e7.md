---
title: Attach a payment to an Invoice
page_id: operation-post-v1-invoices-invoice-attach-payment-788bdbc6
path: operations/untagged
description: |-
    <p>Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of <code>payments</code>.</p>

    <p>For the PaymentIntent, when the PaymentIntent’s status changes to <code>succeeded</code>, the payment is credited
    to the invoice, increasing its <code>amount_paid</code>. When the invoice is fully paid, the
    invoice’s status becomes <code>paid</code>.</p>

    <p>If the PaymentIntent’s status is already <code>succeeded</code> when it’s attached, it’s
    credited to the invoice immediately.</p>

    <p>See: <a href="/docs/invoicing/partial-payments">Partial payments</a> to learn more.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/invoices/{invoice}/attach_payment
operation_ids:
    - PostInvoicesInvoiceAttachPayment
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Attach a payment to an Invoice

`POST /v1/invoices/{invoice}/attach_payment`

Operation ID: `PostInvoicesInvoiceAttachPayment`

<p>Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of <code>payments</code>.</p>

<p>For the PaymentIntent, when the PaymentIntent’s status changes to <code>succeeded</code>, the payment is credited
to the invoice, increasing its <code>amount_paid</code>. When the invoice is fully paid, the
invoice’s status becomes <code>paid</code>.</p>

<p>If the PaymentIntent’s status is already <code>succeeded</code> when it’s attached, it’s
credited to the invoice immediately.</p>

<p>See: <a href="/docs/invoicing/partial-payments">Partial payments</a> to learn more.</p>

## Definition

```yaml
{"summary": "Attach a payment to an Invoice", "description": "<p>Attaches a PaymentIntent or an Out of Band Payment to the invoice, adding it to the list of <code>payments</code>.</p>\n\n<p>For the PaymentIntent, when the PaymentIntent’s status changes to <code>succeeded</code>, the payment is credited\nto the invoice, increasing its <code>amount_paid</code>. When the invoice is fully paid, the\ninvoice’s status becomes <code>paid</code>.</p>\n\n<p>If the PaymentIntent’s status is already <code>succeeded</code> when it’s attached, it’s\ncredited to the invoice immediately.</p>\n\n<p>See: <a href=\"/docs/invoicing/partial-payments\">Partial payments</a> to learn more.</p>", "operationId": "PostInvoicesInvoiceAttachPayment", "parameters": [{"name": "invoice", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "payment_intent": {"maxLength": 5000, "type": "string", "description": "The ID of the PaymentIntent to attach to the invoice."}, "payment_record": {"maxLength": 5000, "type": "string", "description": "The ID of the PaymentRecord to attach to the invoice."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/invoice"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
