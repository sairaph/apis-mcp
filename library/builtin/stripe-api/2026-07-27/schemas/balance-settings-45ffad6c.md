---
title: balance_settings
page_id: schema-balance-settings-45ffad6c
path: schemas
description: Options for customizing account balances and payout settings for a Stripe platform’s connected accounts.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings

Options for customizing account balances and payout settings for a Stripe platform’s connected accounts.

```yaml
{"title": "BalanceSettingsResourceBalanceSettings", "required": ["object", "payments"], "type": "object", "properties": {"object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["balance_settings"]}, "payments": {"$ref": "#/components/schemas/balance_settings_resource_payments"}}, "description": "Options for customizing account balances and payout settings for a Stripe platform’s connected accounts.", "x-expandableFields": ["payments"], "x-resourceId": "balance_settings"}
```
