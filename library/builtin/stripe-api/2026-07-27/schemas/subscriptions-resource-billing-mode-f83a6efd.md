---
title: subscriptions_resource_billing_mode
page_id: schema-subscriptions-resource-billing-mode-f83a6efd
path: schemas
description: The billing mode of the subscription.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_billing_mode

The billing mode of the subscription.

```yaml
{"title": "SubscriptionsResourceBillingMode", "required": ["type"], "type": "object", "properties": {"flexible": {"description": "Configure behavior for flexible billing mode", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/subscriptions_resource_billing_mode_flexible"}]}, "type": {"type": "string", "description": "Controls how prorations and invoices for subscriptions are calculated and orchestrated.", "enum": ["classic", "flexible"]}, "updated_at": {"type": "integer", "description": "Details on when the current billing_mode was adopted.", "format": "unix-time"}}, "description": "The billing mode of the subscription.", "x-expandableFields": ["flexible"]}
```
