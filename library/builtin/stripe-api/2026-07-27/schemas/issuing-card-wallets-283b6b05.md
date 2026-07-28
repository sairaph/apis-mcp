---
title: issuing_card_wallets
page_id: schema-issuing-card-wallets-283b6b05
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# issuing_card_wallets

```yaml
{"title": "IssuingCardWallets", "required": ["apple_pay", "google_pay"], "type": "object", "properties": {"apple_pay": {"$ref": "#/components/schemas/issuing_card_apple_pay"}, "google_pay": {"$ref": "#/components/schemas/issuing_card_google_pay"}, "primary_account_identifier": {"maxLength": 5000, "type": "string", "description": "Unique identifier for a card used with digital wallets", "nullable": true}}, "description": "", "x-expandableFields": ["apple_pay", "google_pay"]}
```
