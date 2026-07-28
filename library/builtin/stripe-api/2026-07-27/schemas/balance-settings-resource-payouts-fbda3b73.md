---
title: balance_settings_resource_payouts
page_id: schema-balance-settings-resource-payouts-fbda3b73
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_payouts

```yaml
{"title": "BalanceSettingsResourcePayouts", "required": ["status"], "type": "object", "properties": {"automatic_transfer_rules_by_currency": {"type": "object", "additionalProperties": {"type": "array", "items": {"$ref": "#/components/schemas/balance_settings_resource_automatic_transfer_rule"}}, "description": "Configures per-currency rules for automatically transferring funds from the payments balance to a FinancialAccount.", "nullable": true}, "minimum_balance_by_currency": {"type": "object", "additionalProperties": {"type": "integer"}, "description": "The minimum balance amount to retain per currency after automatic payouts. Only funds that exceed these amounts are paid out. Learn more about the [minimum balances for automatic payouts](/payouts/minimum-balances-for-automatic-payouts).", "nullable": true}, "schedule": {"description": "Details on when funds from charges are available, and when they are paid out to an external account. See our [Setting Bank and Debit Card Payouts](https://docs.stripe.com/connect/bank-transfers#payout-information) documentation for details.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/balance_settings_resource_payout_schedule"}]}, "statement_descriptor": {"maxLength": 5000, "type": "string", "description": "The text that appears on the bank account statement for payouts. If not set, this defaults to the platform's bank descriptor as set in the Dashboard.", "nullable": true}, "status": {"type": "string", "description": "Whether the funds in this account can be paid out.", "enum": ["disabled", "enabled"]}}, "description": "", "x-expandableFields": ["automatic_transfer_rules_by_currency", "schedule"]}
```
