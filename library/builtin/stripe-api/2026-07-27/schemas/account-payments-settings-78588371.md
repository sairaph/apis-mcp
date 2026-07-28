---
title: account_payments_settings
page_id: schema-account-payments-settings-78588371
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_payments_settings

```yaml
{"title": "AccountPaymentsSettings", "type": "object", "properties": {"statement_descriptor": {"maxLength": 5000, "type": "string", "description": "The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge.", "nullable": true}, "statement_descriptor_kana": {"maxLength": 5000, "type": "string", "description": "The Kana variation of `statement_descriptor` used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).", "nullable": true}, "statement_descriptor_kanji": {"maxLength": 5000, "type": "string", "description": "The Kanji variation of `statement_descriptor` used for charges in Japan. Japanese statement descriptors have [special requirements](https://docs.stripe.com/get-started/account/statement-descriptors#set-japanese-statement-descriptors).", "nullable": true}}, "description": "", "x-expandableFields": []}
```
