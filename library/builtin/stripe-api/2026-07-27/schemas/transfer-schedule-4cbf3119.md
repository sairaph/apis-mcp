---
title: transfer_schedule
page_id: schema-transfer-schedule-4cbf3119
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# transfer_schedule

```yaml
{"title": "TransferSchedule", "required": ["delay_days", "interval"], "type": "object", "properties": {"delay_days": {"type": "integer", "description": "The number of days charges for the account will be held before being paid out."}, "interval": {"maxLength": 5000, "type": "string", "description": "How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`."}, "monthly_anchor": {"type": "integer", "description": "The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months."}, "monthly_payout_days": {"type": "array", "description": "The days of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.", "items": {"type": "integer"}}, "weekly_anchor": {"maxLength": 5000, "type": "string", "description": "The day of the week funds will be paid out, of the style 'monday', 'tuesday', etc. Only shown if `interval` is weekly."}, "weekly_payout_days": {"type": "array", "description": "The days of the week when available funds are paid out, specified as an array, for example, [`monday`, `tuesday`]. Only shown if `interval` is weekly.", "items": {"type": "string", "enum": ["friday", "monday", "thursday", "tuesday", "wednesday"], "x-stripeBypassValidation": true}}}, "description": "", "x-expandableFields": []}
```
