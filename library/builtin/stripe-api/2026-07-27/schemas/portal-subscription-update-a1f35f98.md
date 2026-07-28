---
title: portal_subscription_update
page_id: schema-portal-subscription-update-a1f35f98
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_subscription_update

```yaml
{"title": "PortalSubscriptionUpdate", "required": ["default_allowed_updates", "enabled", "proration_behavior", "schedule_at_period_end", "trial_update_behavior"], "type": "object", "properties": {"billing_cycle_anchor": {"type": "string", "description": "Determines the value to use for the billing cycle anchor on subscription updates. Valid values are `now` or `unchanged`, and the default value is `unchanged`. Setting the value to `now` resets the subscription's billing cycle anchor to the current time (in UTC). For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).", "nullable": true, "enum": ["now", "unchanged"]}, "default_allowed_updates": {"type": "array", "description": "The types of subscription updates that are supported for items listed in the `products` attribute. When empty, subscriptions are not updateable.", "items": {"type": "string", "enum": ["price", "promotion_code", "quantity"]}}, "enabled": {"type": "boolean", "description": "Whether the feature is enabled."}, "products": {"type": "array", "description": "The list of up to 10 products that support subscription updates.", "nullable": true, "items": {"$ref": "#/components/schemas/portal_subscription_update_product"}}, "proration_behavior": {"type": "string", "description": "Determines how to handle prorations resulting from subscription updates. Valid values are `none`, `create_prorations`, and `always_invoice`. Defaults to a value of `none` if you don't set it during creation.", "enum": ["always_invoice", "create_prorations", "none"]}, "schedule_at_period_end": {"$ref": "#/components/schemas/portal_resource_schedule_update_at_period_end"}, "trial_update_behavior": {"type": "string", "description": "Determines how handle updates to trialing subscriptions. Valid values are `end_trial` and `continue_trial`. Defaults to a value of `end_trial` if you don't set it during creation.", "enum": ["continue_trial", "end_trial"]}}, "description": "", "x-expandableFields": ["products", "schedule_at_period_end"]}
```
