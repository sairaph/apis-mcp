---
title: invoices_payments_invoice_payment_associated_payment
page_id: schema-invoices-payments-invoice-payment-associated-payment-c0ac4d12
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# invoices_payments_invoice_payment_associated_payment

```yaml
{"title": "InvoicesPaymentsInvoicePaymentAssociatedPayment", "required": ["type"], "type": "object", "properties": {"charge": {"description": "ID of the successful charge for this payment when `type` is `charge`.Note: charge is only surfaced if the charge object is not associated with a payment intent. If the charge object does have a payment intent, the Invoice Payment surfaces the payment intent instead.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/charge"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/charge"}]}}, "payment_intent": {"description": "ID of the PaymentIntent associated with this payment when `type` is `payment_intent`. Note: This property is only populated for invoices finalized on or after March 15th, 2019.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_intent"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_intent"}]}}, "payment_record": {"description": "ID of the PaymentRecord associated with this payment when `type` is `payment_record`.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/payment_record"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/payment_record"}]}}, "type": {"type": "string", "description": "Type of payment object associated with this invoice payment.", "enum": ["charge", "payment_intent", "payment_record"]}}, "description": "", "x-expandableFields": ["charge", "payment_intent", "payment_record"]}
```
