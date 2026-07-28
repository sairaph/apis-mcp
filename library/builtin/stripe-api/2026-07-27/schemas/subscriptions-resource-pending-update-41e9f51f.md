---
title: subscriptions_resource_pending_update
page_id: schema-subscriptions-resource-pending-update-41e9f51f
path: schemas
description: |-
    Pending Updates store the changes pending from a previous update that will be applied
    to the Subscription upon successful payment.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# subscriptions_resource_pending_update

Pending Updates store the changes pending from a previous update that will be applied
to the Subscription upon successful payment.

```yaml
{"title": "SubscriptionsResourcePendingUpdate", "required": ["expires_at"], "type": "object", "properties": {"billing_cycle_anchor": {"type": "integer", "description": "If the update is applied, determines the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.", "format": "unix-time", "nullable": true}, "discount": {"description": "The pending subscription-level discount that will be applied when the pending update is applied.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/discount"}]}, "discounts": {"type": "array", "description": "The discounts that will be applied to the subscription when the pending update is applied. Use `expand[]=discounts` to expand each discount.", "nullable": true, "items": {"anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/discount"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/discount"}]}}}, "expires_at": {"type": "integer", "description": "The point after which the changes reflected by this update will be discarded and no longer applied.", "format": "unix-time"}, "metadata": {"type": "object", "additionalProperties": {"maxLength": 500, "type": "string"}, "description": "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.", "nullable": true}, "subscription_items": {"type": "array", "description": "List of subscription items, each with an attached plan, that will be set if the update is applied.", "nullable": true, "items": {"$ref": "#/components/schemas/subscription_item"}}, "trial_end": {"type": "integer", "description": "Unix timestamp representing the end of the trial period the customer will get before being charged for the first time, if the update is applied.", "format": "unix-time", "nullable": true}, "trial_from_plan": {"type": "boolean", "description": "Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://docs.stripe.com/billing/subscriptions/trials) to learn more.", "nullable": true}}, "description": "Pending Updates store the changes pending from a previous update that will be applied\nto the Subscription upon successful payment.", "x-expandableFields": ["discount", "discounts", "subscription_items"]}
```
