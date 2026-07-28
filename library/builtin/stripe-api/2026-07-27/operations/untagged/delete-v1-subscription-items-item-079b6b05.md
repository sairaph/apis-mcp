---
title: Delete a subscription item
page_id: operation-delete-v1-subscription-items-item-5e60b64f
path: operations/untagged
description: <p>Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - DELETE
api_endpoints:
    - /v1/subscription_items/{item}
operation_ids:
    - DeleteSubscriptionItemsItem
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Delete a subscription item

`DELETE /v1/subscription_items/{item}`

Operation ID: `DeleteSubscriptionItemsItem`

<p>Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.</p>

## Definition

```yaml
{"summary": "Delete a subscription item", "description": "<p>Deletes an item from the subscription. Removing a subscription item from a subscription will not cancel the subscription.</p>", "operationId": "DeleteSubscriptionItemsItem", "parameters": [{"name": "item", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {"clear_usage": {"type": "boolean", "description": "Delete all usage for the given subscription item. Allowed only when the current plan's `usage_type` is `metered`."}, "payment_behavior": {"type": "string", "description": "Controls how Stripe handles payment when a subscription update requires payment and `collection_method=charge_automatically`.", "enum": ["allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete"]}, "proration_behavior": {"type": "string", "description": "Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.", "enum": ["always_invoice", "create_prorations", "none"]}, "proration_date": {"type": "integer", "description": "If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](/api/invoices/create_preview) endpoint.", "format": "unix-time"}}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/deleted_subscription_item"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
