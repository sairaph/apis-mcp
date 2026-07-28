---
title: subscriptions_resource_billing_schedules
page_id: schema-subscriptions-resource-billing-schedules-f27e0c04
path: schemas
description: Sets the billing schedule for the subscription.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_schedules

Sets the billing schedule for the subscription.

```yaml
{"title": "SubscriptionsResourceBillingSchedules", "required": ["bill_until", "key"], "type": "object", "properties": {"applies_to": {"type": "array", "description": "Specifies which subscription items the billing schedule applies to.", "nullable": true, "items": {"$ref": "#/components/schemas/subscriptions_resource_billing_schedules_applies_to"}}, "bill_until": {"$ref": "#/components/schemas/subscriptions_resource_billing_schedules_bill_until"}, "key": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the billing schedule."}}, "description": "Sets the billing schedule for the subscription.", "x-expandableFields": ["applies_to", "bill_until"]}
```
