---
title: subscriptions_resource_billing_cycle_anchor_config
page_id: schema-subscriptions-resource-billing-cycle-anchor-config-f5d84b54
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_cycle_anchor_config

```yaml
{"title": "SubscriptionsResourceBillingCycleAnchorConfig", "required": ["day_of_month"], "type": "object", "properties": {"day_of_month": {"type": "integer", "description": "The day of the month of the billing_cycle_anchor."}, "hour": {"type": "integer", "description": "The hour of the day of the billing_cycle_anchor.", "nullable": true}, "minute": {"type": "integer", "description": "The minute of the hour of the billing_cycle_anchor.", "nullable": true}, "month": {"type": "integer", "description": "The month to start full cycle billing periods.", "nullable": true}, "second": {"type": "integer", "description": "The second of the minute of the billing_cycle_anchor.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
