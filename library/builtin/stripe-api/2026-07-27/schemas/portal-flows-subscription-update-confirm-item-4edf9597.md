---
title: portal_flows_subscription_update_confirm_item
page_id: schema-portal-flows-subscription-update-confirm-item-4edf9597
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_flows_subscription_update_confirm_item

```yaml
{"title": "PortalFlowsSubscriptionUpdateConfirmItem", "type": "object", "properties": {"id": {"maxLength": 5000, "type": "string", "description": "The ID of the [subscription item](https://docs.stripe.com/api/subscriptions/object#subscription_object-items-data-id) to be updated.", "nullable": true}, "price": {"maxLength": 5000, "type": "string", "description": "The price the customer should subscribe to through this flow. The price must also be included in the configuration's [`features.subscription_update.products`](https://docs.stripe.com/api/customer_portal/configuration#portal_configuration_object-features-subscription_update-products).", "nullable": true}, "quantity": {"type": "integer", "description": "[Quantity](https://docs.stripe.com/subscriptions/quantities) for this item that the customer should subscribe to through this flow."}}, "description": "", "x-expandableFields": []}
```
