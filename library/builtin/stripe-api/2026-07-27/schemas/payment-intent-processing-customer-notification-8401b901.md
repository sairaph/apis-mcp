---
title: payment_intent_processing_customer_notification
page_id: schema-payment-intent-processing-customer-notification-8401b901
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_processing_customer_notification

```yaml
{"title": "PaymentIntentProcessingCustomerNotification", "type": "object", "properties": {"approval_requested": {"type": "boolean", "description": "Whether customer approval has been requested for this payment. For payments greater than INR 15000 or mandate amount, the customer must provide explicit approval of the payment with their bank.", "nullable": true}, "completes_at": {"type": "integer", "description": "If customer approval is required, they need to provide approval before this time.", "format": "unix-time", "nullable": true}}, "description": "", "x-expandableFields": []}
```
