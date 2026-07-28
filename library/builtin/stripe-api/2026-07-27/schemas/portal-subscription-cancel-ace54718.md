---
title: portal_subscription_cancel
page_id: schema-portal-subscription-cancel-ace54718
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_subscription_cancel

```yaml
{"title": "PortalSubscriptionCancel", "required": ["cancellation_reason", "enabled", "mode", "proration_behavior"], "type": "object", "properties": {"cancellation_reason": {"$ref": "#/components/schemas/portal_subscription_cancellation_reason"}, "enabled": {"type": "boolean", "description": "Whether the feature is enabled."}, "mode": {"type": "string", "description": "Whether to cancel subscriptions immediately or at the end of the billing period.", "enum": ["at_period_end", "immediately"]}, "proration_behavior": {"type": "string", "description": "Whether to create prorations when canceling subscriptions. Possible values are `none` and `create_prorations`.", "enum": ["always_invoice", "create_prorations", "none"]}}, "description": "", "x-expandableFields": ["cancellation_reason"]}
```
