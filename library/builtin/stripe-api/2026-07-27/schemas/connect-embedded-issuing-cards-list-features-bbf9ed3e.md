---
title: connect_embedded_issuing_cards_list_features
page_id: schema-connect-embedded-issuing-cards-list-features-bbf9ed3e
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# connect_embedded_issuing_cards_list_features

```yaml
{"title": "ConnectEmbeddedIssuingCardsListFeatures", "required": ["card_management", "card_spend_dispute_management", "cardholder_management", "disable_stripe_user_authentication", "spend_control_management"], "type": "object", "properties": {"card_management": {"type": "boolean", "description": "Whether to allow card management features."}, "card_spend_dispute_management": {"type": "boolean", "description": "Whether to allow card spend dispute management features."}, "cardholder_management": {"type": "boolean", "description": "Whether to allow cardholder management features."}, "disable_stripe_user_authentication": {"type": "boolean", "description": "Whether Stripe user authentication is disabled. This value can only be `true` for accounts where `controller.requirement_collection` is `application` for the account. The default value is the opposite of the `external_account_collection` value. For example, if you don't set `external_account_collection`, it defaults to `true` and `disable_stripe_user_authentication` defaults to `false`."}, "spend_control_management": {"type": "boolean", "description": "Whether to allow spend control management features."}}, "description": "", "x-expandableFields": []}
```
