---
title: mandate_payto
page_id: schema-mandate-payto-2776eec4
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# mandate_payto

```yaml
{"title": "mandate_payto", "required": ["amount_type", "payment_schedule"], "type": "object", "properties": {"amount": {"type": "integer", "description": "Amount that will be collected. It is required when `amount_type` is `fixed`.", "nullable": true}, "amount_type": {"type": "string", "description": "The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.", "enum": ["fixed", "maximum"]}, "end_date": {"maxLength": 5000, "type": "string", "description": "Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.", "nullable": true}, "payment_schedule": {"type": "string", "description": "The periodicity at which payments will be collected. Defaults to `adhoc`.", "enum": ["adhoc", "annual", "daily", "fortnightly", "monthly", "quarterly", "semi_annual", "weekly"]}, "payments_per_period": {"type": "integer", "description": "The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.", "nullable": true}, "purpose": {"type": "string", "description": "The purpose for which payments are made. Has a default value based on your merchant category code.", "nullable": true, "enum": ["dependant_support", "government", "loan", "mortgage", "other", "pension", "personal", "retail", "salary", "tax", "utility"]}, "start_date": {"maxLength": 5000, "type": "string", "description": "Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
