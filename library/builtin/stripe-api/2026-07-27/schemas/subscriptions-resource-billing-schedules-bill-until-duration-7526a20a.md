---
title: subscriptions_resource_billing_schedules_bill_until_duration
page_id: schema-subscriptions-resource-billing-schedules-bill-until-duration-7526a20a
path: schemas
description: Configures the `bill_until` date based on the provided `interval` and `interval_count`.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_schedules_bill_until_duration

Configures the `bill_until` date based on the provided `interval` and `interval_count`.

```yaml
{"title": "SubscriptionsResourceBillingSchedulesBillUntilDuration", "required": ["interval"], "type": "object", "properties": {"interval": {"type": "string", "description": "Specifies billing duration. Either `day`, `week`, `month` or `year`.", "enum": ["day", "month", "week", "year"]}, "interval_count": {"type": "integer", "description": "The multiplier applied to the interval.", "nullable": true}}, "description": "Configures the `bill_until` date based on the provided `interval` and `interval_count`.", "x-expandableFields": []}
```
