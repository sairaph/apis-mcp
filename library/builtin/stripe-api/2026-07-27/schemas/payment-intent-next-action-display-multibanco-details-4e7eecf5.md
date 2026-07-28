---
title: payment_intent_next_action_display_multibanco_details
page_id: schema-payment-intent-next-action-display-multibanco-details-4e7eecf5
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_display_multibanco_details

```yaml
{"title": "PaymentIntentNextActionDisplayMultibancoDetails", "type": "object", "properties": {"entity": {"maxLength": 5000, "type": "string", "description": "Entity number associated with this Multibanco payment.", "nullable": true}, "expires_at": {"type": "integer", "description": "The timestamp at which the Multibanco voucher expires.", "format": "unix-time", "nullable": true}, "hosted_voucher_url": {"maxLength": 5000, "type": "string", "description": "The URL for the hosted Multibanco voucher page, which allows customers to view a Multibanco voucher.", "nullable": true}, "reference": {"maxLength": 5000, "type": "string", "description": "Reference number associated with this Multibanco payment.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
