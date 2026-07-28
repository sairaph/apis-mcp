---
title: balance_settings_resource_settlement_timing
page_id: schema-balance-settings-resource-settlement-timing-be45db18
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_settlement_timing

```yaml
{"title": "BalanceSettingsResourceSettlementTiming", "required": ["delay_days"], "type": "object", "properties": {"delay_days": {"type": "integer", "description": "The number of days charge funds are held before becoming available."}, "delay_days_override": {"type": "integer", "description": "The number of days charge funds are held before becoming available. If present, overrides the default, or minimum available, for the account."}, "start_of_day": {"description": "Customized start of day configuration for automatic payouts to group and send payments in local timezones with a customized day starting time. For details, see our [Customized start of day](/connect/customized-start-of-day) documentation.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/balance_settings_resource_start_of_day"}]}}, "description": "", "x-expandableFields": ["start_of_day"]}
```
