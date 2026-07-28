---
title: subscription_schedule_current_phase
page_id: schema-subscription-schedule-current-phase-40514f0f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_schedule_current_phase

```yaml
{"title": "SubscriptionScheduleCurrentPhase", "required": ["end_date", "start_date"], "type": "object", "properties": {"end_date": {"type": "integer", "description": "The end of this phase of the subscription schedule.", "format": "unix-time"}, "start_date": {"type": "integer", "description": "The start of this phase of the subscription schedule.", "format": "unix-time"}}, "description": "", "x-expandableFields": []}
```
