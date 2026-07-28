---
title: Pay an invoice
page_id: operation-post-v1-invoices-invoice-pay-12e1fad6
path: operations/untagged
description: <p>Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your <a href="https://dashboard.stripe.com/account/billing/automatic">subscriptions settings</a>. However, if you’d like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - POST
api_endpoints:
    - /v1/invoices/{invoice}/pay
operation_ids:
    - PostInvoicesInvoicePay
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Pay an invoice

`POST /v1/invoices/{invoice}/pay`

Operation ID: `PostInvoicesInvoicePay`

<p>Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your <a href="https://dashboard.stripe.com/account/billing/automatic">subscriptions settings</a>. However, if you’d like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.</p>

## Definition

```yaml
{"summary": "Pay an invoice", "description": "<p>Stripe automatically creates and then attempts to collect payment on invoices for customers on subscriptions according to your <a href=\"https://dashboard.stripe.com/account/billing/automatic\">subscriptions settings</a>. However, if you’d like to attempt payment on an invoice out of the normal collection schedule or for some other reason, you can do so.</p>", "operationId": "PostInvoicesInvoicePay", "parameters": [{"name": "invoice", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"expand": {"type": "array", "description": "Specifies which fields in the response should be expanded.", "items": {"maxLength": 5000, "type": "string"}}, "forgive": {"type": "boolean", "description": "In cases where the source used to pay the invoice has insufficient funds, passing `forgive=true` controls whether a charge should be attempted for the full amount available on the source, up to the amount to fully pay the invoice. This effectively forgives the difference between the amount available on the source and the amount due. \n\nPassing `forgive=false` will fail the charge if the source hasn't been pre-funded with the right amount. An example for this case is with ACH Credit Transfers and wires: if the amount wired is less than the amount due by a small amount, you might want to forgive the difference. Defaults to `false`."}, "mandate": {"description": "ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the payment_method param or the invoice's default_payment_method or default_source, if set.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"type": "string", "enum": [""]}]}, "off_session": {"type": "boolean", "description": "Indicates if a customer is on or off-session while an invoice payment is attempted. Defaults to `true` (off-session)."}, "paid_out_of_band": {"type": "boolean", "description": "Boolean representing whether an invoice is paid outside of Stripe. This will result in no charge being made. Defaults to `false`."}, "payment_method": {"maxLength": 5000, "type": "string", "description": "A PaymentMethod to be charged. The PaymentMethod must be the ID of a PaymentMethod belonging to the customer associated with the invoice being paid."}, "source": {"maxLength": 5000, "type": "string", "description": "A payment source to be charged. The source must be the ID of a source belonging to the customer associated with the invoice being paid."}}, "additionalProperties": false}, "encoding": {"expand": {"style": "deepObject", "explode": true}, "mandate": {"style": "deepObject", "explode": true}}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/invoice"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
