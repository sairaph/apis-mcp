---
title: payment_intent_next_action_display_oxxo_details
page_id: schema-payment-intent-next-action-display-oxxo-details-c0131230
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_display_oxxo_details

```yaml
{"title": "PaymentIntentNextActionDisplayOxxoDetails", "type": "object", "properties": {"expires_after": {"type": "integer", "description": "The timestamp after which the OXXO voucher expires.", "format": "unix-time", "nullable": true}, "hosted_voucher_url": {"maxLength": 5000, "type": "string", "description": "The URL for the hosted OXXO voucher page, which allows customers to view and print an OXXO voucher.", "nullable": true}, "number": {"maxLength": 5000, "type": "string", "description": "OXXO reference number.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
