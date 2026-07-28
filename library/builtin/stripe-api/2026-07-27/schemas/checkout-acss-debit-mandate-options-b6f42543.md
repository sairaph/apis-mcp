---
title: checkout_acss_debit_mandate_options
page_id: schema-checkout-acss-debit-mandate-options-b6f42543
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# checkout_acss_debit_mandate_options

```yaml
{"title": "CheckoutAcssDebitMandateOptions", "type": "object", "properties": {"custom_mandate_url": {"maxLength": 5000, "type": "string", "description": "A URL for custom mandate text"}, "default_for": {"type": "array", "description": "List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.", "items": {"type": "string", "enum": ["invoice", "subscription"]}}, "interval_description": {"maxLength": 5000, "type": "string", "description": "Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.", "nullable": true}, "payment_schedule": {"type": "string", "description": "Payment schedule for the mandate.", "nullable": true, "enum": ["combined", "interval", "sporadic"]}, "transaction_type": {"type": "string", "description": "Transaction type of the mandate.", "nullable": true, "enum": ["business", "personal"]}}, "description": "", "x-expandableFields": []}
```
