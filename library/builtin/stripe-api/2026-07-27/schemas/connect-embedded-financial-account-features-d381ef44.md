---
title: connect_embedded_financial_account_features
page_id: schema-connect-embedded-financial-account-features-d381ef44
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_financial_account_features

```yaml
{"title": "ConnectEmbeddedFinancialAccountFeatures", "required": ["disable_stripe_user_authentication", "external_account_collection", "send_money", "transfer_balance"], "type": "object", "properties": {"disable_stripe_user_authentication": {"type": "boolean", "description": "Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`."}, "external_account_collection": {"type": "boolean", "description": "Whether external account collection is enabled. This feature can only be `false` for accounts where you’re responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`."}, "send_money": {"type": "boolean", "description": "Whether to allow sending money."}, "transfer_balance": {"type": "boolean", "description": "Whether to allow transferring balance."}}, "description": "", "x-expandableFields": []}
```
