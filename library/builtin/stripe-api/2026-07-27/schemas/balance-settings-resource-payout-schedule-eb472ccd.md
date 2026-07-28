---
title: balance_settings_resource_payout_schedule
page_id: schema-balance-settings-resource-payout-schedule-eb472ccd
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_payout_schedule

```yaml
{"title": "BalanceSettingsResourcePayoutSchedule", "type": "object", "properties": {"interval": {"type": "string", "description": "How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.", "nullable": true, "enum": ["daily", "manual", "monthly", "weekly"]}, "monthly_payout_days": {"type": "array", "description": "The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.", "items": {"type": "integer"}}, "weekly_payout_days": {"type": "array", "description": "The days of the week when available funds are paid out, specified as an array, for example, [`monday`, `tuesday`]. Only shown if `interval` is weekly.", "items": {"type": "string", "enum": ["friday", "monday", "thursday", "tuesday", "wednesday"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": []}
```
