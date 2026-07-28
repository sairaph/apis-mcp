---
title: connect_embedded_account_features_claim
page_id: schema-connect-embedded-account-features-claim-54f0282f
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_account_features_claim

```yaml
{"title": "ConnectEmbeddedAccountFeaturesClaim", "required": ["disable_stripe_user_authentication", "external_account_collection"], "type": "object", "properties": {"disable_stripe_user_authentication": {"type": "boolean", "description": "Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`."}, "external_account_collection": {"type": "boolean", "description": "Whether external account collection is enabled. This feature can only be `false` for accounts where you’re responsible for collecting updated information when requirements are due or change, like Custom accounts. The default value for this feature is `true`."}}, "description": "", "x-expandableFields": []}
```
