---
title: payment_intent_processing
page_id: schema-payment-intent-processing-8885d871
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_processing

```yaml
{"title": "PaymentIntentProcessing", "required": ["type"], "type": "object", "properties": {"card": {"$ref": "#/components/schemas/payment_intent_card_processing"}, "type": {"type": "string", "description": "Type of the payment method for which payment is in `processing` state, one of `card`.", "enum": ["card"]}}, "description": "", "x-expandableFields": ["card"]}
```
