---
title: balance_settings_resource_automatic_transfer_rule
page_id: schema-balance-settings-resource-automatic-transfer-rule-ca564072
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_automatic_transfer_rule

```yaml
{"title": "BalanceSettingsResourceAutomaticTransferRule", "required": ["payout_method", "type"], "type": "object", "properties": {"payout_method": {"maxLength": 5000, "type": "string", "description": "The ID of the FinancialAccount that funds will be transferred to during automatic transfers."}, "transfer_up_to_amount": {"type": "integer", "description": "The maximum amount in minor units to transfer to the FinancialAccount. Only applicable when `type` is `transfer_up_to_amount`.", "nullable": true}, "type": {"type": "string", "description": "The type of automatic transfer rule.", "enum": ["transfer_all", "transfer_up_to_amount"]}}, "description": "", "x-expandableFields": []}
```
