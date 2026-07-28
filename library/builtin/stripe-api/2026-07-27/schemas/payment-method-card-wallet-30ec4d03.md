---
title: payment_method_card_wallet
page_id: schema-payment-method-card-wallet-30ec4d03
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_card_wallet

```yaml
{"title": "payment_method_card_wallet", "required": ["type"], "type": "object", "properties": {"amex_express_checkout": {"$ref": "#/components/schemas/payment_method_card_wallet_amex_express_checkout"}, "apple_pay": {"$ref": "#/components/schemas/payment_method_card_wallet_apple_pay"}, "dynamic_last4": {"maxLength": 5000, "type": "string", "description": "(For tokenized numbers only.) The last four digits of the device account number.", "nullable": true}, "google_pay": {"$ref": "#/components/schemas/payment_method_card_wallet_google_pay"}, "link": {"$ref": "#/components/schemas/payment_method_card_wallet_link"}, "masterpass": {"$ref": "#/components/schemas/payment_method_card_wallet_masterpass"}, "samsung_pay": {"$ref": "#/components/schemas/payment_method_card_wallet_samsung_pay"}, "type": {"type": "string", "description": "The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.", "enum": ["amex_express_checkout", "apple_pay", "google_pay", "link", "masterpass", "samsung_pay", "visa_checkout"]}, "visa_checkout": {"$ref": "#/components/schemas/payment_method_card_wallet_visa_checkout"}}, "description": "", "x-expandableFields": ["amex_express_checkout", "apple_pay", "google_pay", "link", "masterpass", "samsung_pay", "visa_checkout"]}
```
