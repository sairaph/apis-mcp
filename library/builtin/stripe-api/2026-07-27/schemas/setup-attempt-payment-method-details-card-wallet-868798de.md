---
title: setup_attempt_payment_method_details_card_wallet
page_id: schema-setup-attempt-payment-method-details-card-wallet-868798de
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# setup_attempt_payment_method_details_card_wallet

```yaml
{"title": "setup_attempt_payment_method_details_card_wallet", "required": ["type"], "type": "object", "properties": {"apple_pay": {"$ref": "#/components/schemas/payment_method_details_card_wallet_apple_pay"}, "google_pay": {"$ref": "#/components/schemas/payment_method_details_card_wallet_google_pay"}, "type": {"type": "string", "description": "The type of the card wallet, one of `apple_pay`, `google_pay`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.", "enum": ["apple_pay", "google_pay", "link"]}}, "description": "", "x-expandableFields": ["apple_pay", "google_pay"]}
```
