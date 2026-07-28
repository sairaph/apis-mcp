---
title: subscriptions_resource_billing_schedules_bill_until
page_id: schema-subscriptions-resource-billing-schedules-bill-until-ff93d5cc
path: schemas
description: Specifies the end of billing period.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_schedules_bill_until

Specifies the end of billing period.

```yaml
{"title": "SubscriptionsResourceBillingSchedulesBillUntil", "required": ["computed_timestamp", "type"], "type": "object", "properties": {"computed_timestamp": {"type": "integer", "description": "The timestamp the billing schedule will apply until.", "format": "unix-time"}, "duration": {"description": "Specifies the billing period.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscriptions_resource_billing_schedules_bill_until_duration"}]}, "timestamp": {"type": "integer", "description": "If specified, the billing schedule will apply until the specified timestamp.", "format": "unix-time", "nullable": true}, "type": {"type": "string", "description": "Describes how the billing schedule will determine the end date. Either `duration` or `timestamp`.", "enum": ["duration", "timestamp"]}}, "description": "Specifies the end of billing period.", "x-expandableFields": ["duration"]}
```
