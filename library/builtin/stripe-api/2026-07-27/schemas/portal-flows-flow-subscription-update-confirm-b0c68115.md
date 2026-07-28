---
title: portal_flows_flow_subscription_update_confirm
page_id: schema-portal-flows-flow-subscription-update-confirm-b0c68115
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_flow_subscription_update_confirm

```yaml
{"title": "PortalFlowsFlowSubscriptionUpdateConfirm", "required": ["items", "subscription"], "type": "object", "properties": {"discounts": {"type": "array", "description": "The coupon or promotion code to apply to this subscription update.", "nullable": true, "items": {"$ref": "#/components/schemas/portal_flows_subscription_update_confirm_discount"}}, "items": {"type": "array", "description": "The [subscription item](https://docs.stripe.com/api/subscription_items) to be updated through this flow. Currently, only up to one may be specified and subscriptions with multiple items are not updatable.", "items": {"$ref": "#/components/schemas/portal_flows_subscription_update_confirm_item"}}, "subscription": {"maxLength": 5000, "type": "string", "description": "The ID of the subscription to be updated."}}, "description": "", "x-expandableFields": ["discounts", "items"]}
```
