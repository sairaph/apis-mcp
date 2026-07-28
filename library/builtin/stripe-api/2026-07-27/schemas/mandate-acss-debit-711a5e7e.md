---
title: mandate_acss_debit
page_id: schema-mandate-acss-debit-711a5e7e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_acss_debit

```yaml
{"title": "mandate_acss_debit", "required": ["payment_schedule", "transaction_type"], "type": "object", "properties": {"default_for": {"type": "array", "description": "List of Stripe products where this mandate can be selected automatically.", "items": {"type": "string", "enum": ["invoice", "subscription"]}}, "interval_description": {"maxLength": 5000, "type": "string", "description": "Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.", "nullable": true}, "payment_schedule": {"type": "string", "description": "Payment schedule for the mandate.", "enum": ["combined", "interval", "sporadic"]}, "transaction_type": {"type": "string", "description": "Transaction type of the mandate.", "enum": ["business", "personal"]}}, "description": "", "x-expandableFields": []}
```
