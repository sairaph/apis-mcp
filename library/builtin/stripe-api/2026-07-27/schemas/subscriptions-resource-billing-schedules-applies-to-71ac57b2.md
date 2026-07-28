---
title: subscriptions_resource_billing_schedules_applies_to
page_id: schema-subscriptions-resource-billing-schedules-applies-to-71ac57b2
path: schemas
description: Represents the entities that the billing schedule applies to.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_schedules_applies_to

Represents the entities that the billing schedule applies to.

```yaml
{"title": "SubscriptionsResourceBillingSchedulesAppliesTo", "required": ["type"], "type": "object", "properties": {"price": {"description": "The billing schedule will apply to the subscription item with the given price ID.", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/price"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/price"}]}}, "type": {"type": "string", "description": "Controls which subscription items the billing schedule applies to.", "enum": ["price"]}}, "description": "Represents the entities that the billing schedule applies to.", "x-expandableFields": ["price"]}
```
