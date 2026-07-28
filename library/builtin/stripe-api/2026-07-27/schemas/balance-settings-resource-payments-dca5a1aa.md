---
title: balance_settings_resource_payments
page_id: schema-balance-settings-resource-payments-dca5a1aa
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_payments

```yaml
{"title": "BalanceSettingsResourcePayments", "required": ["settlement_timing"], "type": "object", "properties": {"debit_negative_balances": {"type": "boolean", "description": "A Boolean indicating if Stripe should try to reclaim negative balances from an attached bank account. See [Understanding Connect account balances](/connect/account-balances) for details. The default value is `false` when [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `application`, which includes Custom accounts, otherwise `true`.", "nullable": true}, "payouts": {"description": "Settings specific to the account's payouts.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/balance_settings_resource_payouts"}]}, "settlement_timing": {"$ref": "#/components/schemas/balance_settings_resource_settlement_timing"}}, "description": "", "x-expandableFields": ["payouts", "settlement_timing"]}
```
