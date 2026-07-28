---
title: invoice_payment
page_id: schema-invoice-payment-159da092
path: schemas
description: |-
    Invoice Payments represent payments made against invoices. Invoice Payments can
    be accessed in two ways:
    1. By expanding the `payments` field on the [Invoice](https://api.stripe.com#invoice) resource.
    2. By using the Invoice Payment retrieve and list endpoints.

    Invoice Payments include the mapping between payment objects, such as Payment Intent, and Invoices.
    This resource and its endpoints allows you to easily track if a payment is associated with a specific invoice and
    monitor the allocation details of the payments.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoice_payment

Invoice Payments represent payments made against invoices. Invoice Payments can
be accessed in two ways:
1. By expanding the `payments` field on the [Invoice](https://api.stripe.com#invoice) resource.
2. By using the Invoice Payment retrieve and list endpoints.

Invoice Payments include the mapping between payment objects, such as Payment Intent, and Invoices.
This resource and its endpoints allows you to easily track if a payment is associated with a specific invoice and
monitor the allocation details of the payments.

```yaml
{"title": "InvoicesInvoicePayment", "required": ["amount_requested", "created", "currency", "id", "invoice", "is_default", "livemode", "object", "payment", "status", "status_transitions"], "type": "object", "properties": {"amount_paid": {"type": "integer", "description": "Amount that was actually paid for this invoice, in cents (or local equivalent). This field is null until the payment is `paid`. This amount can be less than the `amount_requested` if the PaymentIntent’s `amount_received` is not sufficient to pay all of the invoices that it is attached to.", "nullable": true}, "amount_requested": {"type": "integer", "description": "Amount intended to be paid toward this invoice, in cents (or local equivalent)"}, "created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "currency": {"maxLength": 5000, "type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies)."}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "invoice": {"description": "The invoice that was paid.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/invoice"}, {"$ref": "#/components/schemas/deleted_invoice"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/invoice"}, {"$ref": "#/components/schemas/deleted_invoice"}]}}, "is_default": {"type": "boolean", "description": "Stripe automatically creates a default InvoicePayment when the invoice is finalized, and keeps it synchronized with the invoice’s `amount_remaining`. The PaymentIntent associated with the default payment can’t be edited or canceled directly."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["invoice_payment"]}, "payment": {"$ref": "#/components/schemas/invoices_payments_invoice_payment_associated_payment"}, "status": {"maxLength": 5000, "type": "string", "description": "The status of the payment, one of `open`, `paid`, or `canceled`."}, "status_transitions": {"$ref": "#/components/schemas/invoices_payments_invoice_payment_status_transitions"}}, "description": "Invoice Payments represent payments made against invoices. Invoice Payments can\nbe accessed in two ways:\n1. By expanding the `payments` field on the [Invoice](https://api.stripe.com#invoice) resource.\n2. By using the Invoice Payment retrieve and list endpoints.\n\nInvoice Payments include the mapping between payment objects, such as Payment Intent, and Invoices.\nThis resource and its endpoints allows you to easily track if a payment is associated with a specific invoice and\nmonitor the allocation details of the payments.", "x-expandableFields": ["invoice", "payment", "status_transitions"], "x-resourceId": "invoice_payment"}
```
