---
title: dispute_payment_method_details_klarna
page_id: schema-dispute-payment-method-details-klarna-2e79ed33
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# dispute_payment_method_details_klarna

```yaml
{"title": "DisputePaymentMethodDetailsKlarna", "type": "object", "properties": {"chargeback_loss_reason_code": {"maxLength": 5000, "type": "string", "description": "Chargeback loss reason mapped by Stripe from Klarna's chargeback loss reason"}, "reason_code": {"maxLength": 5000, "type": "string", "description": "The reason for the dispute as defined by Klarna", "nullable": true}}, "description": "", "x-expandableFields": []}
```
