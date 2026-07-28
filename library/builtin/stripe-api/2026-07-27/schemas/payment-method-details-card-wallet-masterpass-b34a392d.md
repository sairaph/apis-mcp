---
title: payment_method_details_card_wallet_masterpass
page_id: schema-payment-method-details-card-wallet-masterpass-b34a392d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payment_method_details_card_wallet_masterpass

```yaml
{"title": "payment_method_details_card_wallet_masterpass", "type": "object", "properties": {"billing_address": {"description": "Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}, "email": {"maxLength": 5000, "type": "string", "description": "Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "name": {"maxLength": 5000, "type": "string", "description": "Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true}, "shipping_address": {"description": "Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/address"}]}}, "description": "", "x-expandableFields": ["billing_address", "shipping_address"]}
```
