---
title: payment_flows_payment_intent_presentment_details
page_id: schema-payment-flows-payment-intent-presentment-details-3bb64f92
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_flows_payment_intent_presentment_details

```yaml
{"title": "PaymentFlowsPaymentIntentPresentmentDetails", "required": ["presentment_amount", "presentment_currency"], "type": "object", "properties": {"presentment_amount": {"type": "integer", "description": "Amount intended to be collected by this payment, denominated in `presentment_currency`."}, "presentment_currency": {"maxLength": 5000, "type": "string", "description": "Currency presented to the customer during payment."}}, "description": "", "x-expandableFields": []}
```
