---
title: payments_primitives_payment_records_resource_payment_method_card_details_resource_wallet
page_id: schema-payments-primitives-payment-records-resource-payment-method-card-details-cccd14fb
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payments_primitives_payment_records_resource_payment_method_card_details_resource_wallet

```yaml
{"title": "PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCardDetailsResourceWallet", "required": ["type"], "type": "object", "properties": {"apple_pay": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_payment_method_card_details_resource_wallet_resource_apple_pay"}, "dynamic_last4": {"maxLength": 5000, "type": "string", "description": "(For tokenized numbers only.) The last four digits of the device account number."}, "google_pay": {"$ref": "#/components/schemas/payments_primitives_payment_records_resource_payment_method_card_details_resource_wallet_resource_google_pay"}, "type": {"maxLength": 5000, "type": "string", "description": "The type of the card wallet, one of `apple_pay` or `google_pay`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type."}}, "description": "", "x-expandableFields": ["apple_pay", "google_pay"]}
```
