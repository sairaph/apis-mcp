---
title: payment_method_details_card_present_offline
page_id: schema-payment-method-details-card-present-offline-b073844b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_card_present_offline

```yaml
{"title": "payment_method_details_card_present_offline", "type": "object", "properties": {"stored_at": {"type": "integer", "description": "Time at which the payment was collected while offline", "format": "unix-time", "nullable": true}, "type": {"type": "string", "description": "The method used to process this payment method offline. Only deferred is allowed.", "nullable": true, "enum": ["deferred"]}}, "description": "", "x-expandableFields": []}
```
