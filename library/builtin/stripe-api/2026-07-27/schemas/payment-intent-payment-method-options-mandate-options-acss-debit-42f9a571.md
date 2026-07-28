---
title: payment_intent_payment_method_options_mandate_options_acss_debit
page_id: schema-payment-intent-payment-method-options-mandate-options-acss-debit-42f9a571
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_payment_method_options_mandate_options_acss_debit

```yaml
{"title": "payment_intent_payment_method_options_mandate_options_acss_debit", "type": "object", "properties": {"custom_mandate_url": {"maxLength": 5000, "type": "string", "description": "A URL for custom mandate text"}, "interval_description": {"maxLength": 5000, "type": "string", "description": "Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.", "nullable": true}, "payment_schedule": {"type": "string", "description": "Payment schedule for the mandate.", "nullable": true, "enum": ["combined", "interval", "sporadic"]}, "transaction_type": {"type": "string", "description": "Transaction type of the mandate.", "nullable": true, "enum": ["business", "personal"]}}, "description": "", "x-expandableFields": []}
```
