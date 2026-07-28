---
title: payment_method_options_mandate_options_pix
page_id: schema-payment-method-options-mandate-options-pix-e511c4b8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_options_mandate_options_pix

```yaml
{"title": "payment_method_options_mandate_options_pix", "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount to be charged for future payments."}, "amount_includes_iof": {"type": "string", "description": "Determines if the amount includes the IOF tax.", "enum": ["always", "never"]}, "amount_type": {"type": "string", "description": "Type of amount.", "enum": ["fixed", "maximum"]}, "currency": {"type": "string", "description": "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.", "format": "currency"}, "end_date": {"maxLength": 5000, "type": "string", "description": "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`."}, "payment_schedule": {"type": "string", "description": "Schedule at which the future payments will be charged.", "enum": ["halfyearly", "monthly", "quarterly", "weekly", "yearly"]}, "reference": {"maxLength": 5000, "type": "string", "description": "Subscription name displayed to buyers in their bank app."}, "start_date": {"maxLength": 5000, "type": "string", "description": "Start date of the mandate, in `YYYY-MM-DD`."}}, "description": "", "x-expandableFields": []}
```
