---
title: balance
page_id: schema-balance-5751e048
path: schemas
description: |-
    This is an object representing your Stripe balance. You can retrieve it to see
    the balance currently on your Stripe account.

    The top-level `available` and `pending` comprise your "payments balance."

    Related guide: [Balances and settlement time](https://docs.stripe.com/payments/balances), [Understanding Connect account balances](https://docs.stripe.com/connect/account-balances)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance

This is an object representing your Stripe balance. You can retrieve it to see
the balance currently on your Stripe account.

The top-level `available` and `pending` comprise your "payments balance."

Related guide: [Balances and settlement time](https://docs.stripe.com/payments/balances), [Understanding Connect account balances](https://docs.stripe.com/connect/account-balances)

```yaml
{"title": "Balance", "required": ["available", "livemode", "object", "pending"], "type": "object", "properties": {"available": {"type": "array", "description": "Available funds that you can transfer or pay out automatically by Stripe or explicitly through the [Transfers API](https://api.stripe.com#transfers) or [Payouts API](https://api.stripe.com#payouts). You can find the available balance for each currency and payment type in the `source_types` property.", "items": {"$ref": "#/components/schemas/balance_amount"}}, "connect_reserved": {"type": "array", "description": "Funds held due to negative balances on connected accounts where [account.controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts. You can find the connect reserve balance for each currency and payment type in the `source_types` property.", "items": {"$ref": "#/components/schemas/balance_amount"}}, "instant_available": {"type": "array", "description": "Funds that you can pay out using Instant Payouts.", "items": {"$ref": "#/components/schemas/balance_amount_net"}}, "issuing": {"$ref": "#/components/schemas/balance_detail"}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["balance"]}, "pending": {"type": "array", "description": "Funds that aren't available in the balance yet. You can find the pending balance for each currency and each payment type in the `source_types` property.", "items": {"$ref": "#/components/schemas/balance_amount"}}, "refund_and_dispute_prefunding": {"$ref": "#/components/schemas/balance_detail_ungated"}}, "description": "This is an object representing your Stripe balance. You can retrieve it to see\nthe balance currently on your Stripe account.\n\nThe top-level `available` and `pending` comprise your \"payments balance.\"\n\nRelated guide: [Balances and settlement time](https://docs.stripe.com/payments/balances), [Understanding Connect account balances](https://docs.stripe.com/connect/account-balances)", "x-expandableFields": ["available", "connect_reserved", "instant_available", "issuing", "pending", "refund_and_dispute_prefunding"], "x-resourceId": "balance"}
```
