---
title: subscription_payment_method_options_mandate_options_pix
page_id: schema-subscription-payment-method-options-mandate-options-pix-cbe8bd63
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_payment_method_options_mandate_options_pix

```yaml
{"title": "subscription_payment_method_options_mandate_options_pix", "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be charged for future payments.", "nullable": true}, "amount_includes_iof": {"type": "string", "description": "Determines if the amount includes the IOF tax.", "nullable": true, "enum": ["always", "never"]}, "end_date": {"maxLength": 5000, "type": "string", "description": "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`.", "nullable": true}, "payment_schedule": {"type": "string", "description": "Schedule at which the future payments will be charged.", "nullable": true, "enum": ["halfyearly", "monthly", "quarterly", "weekly", "yearly"]}}, "description": "", "x-expandableFields": []}
```
