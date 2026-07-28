---
title: connect_embedded_payouts_features
page_id: schema-connect-embedded-payouts-features-2eb89f44
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_payouts_features

```yaml
{"title": "ConnectEmbeddedPayoutsFeatures", "required": ["disable_stripe_user_authentication", "edit_payout_schedule", "external_account_collection", "instant_payouts", "standard_payouts"], "type": "object", "properties": {"disable_stripe_user_authentication": {"type": "boolean", "description": "Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`."}, "edit_payout_schedule": {"type": "boolean", "description": "Whether to allow payout schedule to be changed. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`."}, "external_account_collection": {"type": "boolean", "description": "Whether external account collection is enabled. This feature can only be `false` for accounts where you’re responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`."}, "instant_payouts": {"type": "boolean", "description": "Whether to allow creation of instant payouts. The default value is `enabled` when Stripe is responsible for negative account balances, and `use_dashboard_rules` otherwise."}, "standard_payouts": {"type": "boolean", "description": "Whether to allow creation of standard payouts. Defaults to `true` when `controller.losses.payments` is set to `stripe` for the account, otherwise `false`."}}, "description": "", "x-expandableFields": []}
```
