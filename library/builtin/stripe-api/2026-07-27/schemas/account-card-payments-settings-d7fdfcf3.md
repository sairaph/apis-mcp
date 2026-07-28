---
title: account_card_payments_settings
page_id: schema-account-card-payments-settings-d7fdfcf3
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# account_card_payments_settings

```yaml
{"title": "AccountCardPaymentsSettings", "type": "object", "properties": {"decline_on": {"$ref": "#/components/schemas/account_decline_charge_on"}, "statement_descriptor_prefix": {"maxLength": 5000, "type": "string", "description": "The default text that appears on credit card statements when a charge is made. This field prefixes any dynamic `statement_descriptor` specified on the charge. `statement_descriptor_prefix` is useful for maximizing descriptor space for the dynamic portion.", "nullable": true}, "statement_descriptor_prefix_kana": {"maxLength": 5000, "type": "string", "description": "The Kana variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kana` specified on the charge. `statement_descriptor_prefix_kana` is useful for maximizing descriptor space for the dynamic portion.", "nullable": true}, "statement_descriptor_prefix_kanji": {"maxLength": 5000, "type": "string", "description": "The Kanji variation of the default text that appears on credit card statements when a charge is made (Japan only). This field prefixes any dynamic `statement_descriptor_suffix_kanji` specified on the charge. `statement_descriptor_prefix_kanji` is useful for maximizing descriptor space for the dynamic portion.", "nullable": true}}, "description": "", "x-expandableFields": ["decline_on"]}
```
