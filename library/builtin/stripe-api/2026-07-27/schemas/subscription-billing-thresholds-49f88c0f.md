---
title: subscription_billing_thresholds
page_id: schema-subscription-billing-thresholds-49f88c0f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscription_billing_thresholds

```yaml
{"title": "SubscriptionBillingThresholds", "type": "object", "properties": {"amount_gte": {"type": "integer", "description": "Monetary threshold that triggers the subscription to create an invoice", "nullable": true}, "reset_billing_cycle_anchor": {"type": "boolean", "description": "Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
