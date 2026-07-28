---
title: payment_intent_next_action_card_await_notification
page_id: schema-payment-intent-next-action-card-await-notification-8d3f55b8
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_intent_next_action_card_await_notification

```yaml
{"title": "PaymentIntentNextActionCardAwaitNotification", "type": "object", "properties": {"charge_attempt_at": {"type": "integer", "description": "The time that payment will be attempted. If customer approval is required, they need to provide approval before this time.", "format": "unix-time", "nullable": true}, "customer_approval_required": {"type": "boolean", "description": "For payments greater than INR 15000, the customer must provide explicit approval of the payment with their bank. For payments of lower amount, no customer action is required.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
