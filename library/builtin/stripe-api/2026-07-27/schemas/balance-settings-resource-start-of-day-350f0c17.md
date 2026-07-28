---
title: balance_settings_resource_start_of_day
page_id: schema-balance-settings-resource-start-of-day-350f0c17
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# balance_settings_resource_start_of_day

```yaml
{"title": "BalanceSettingsResourceStartOfDay", "required": ["hour", "minutes", "timezone"], "type": "object", "properties": {"hour": {"type": "integer", "description": "Hour at which the customized start of day begins according to the given timezone. Must be a [supported customized start of day hour](/connect/customized-start-of-day#available-timezones-and-cutoffs)."}, "minutes": {"type": "integer", "description": "Minutes at which the customized start of day begins according to the given timezone. Must be either 0 or 30."}, "timezone": {"maxLength": 5000, "type": "string", "description": "Timezone for the customized start of day. Must be a [supported customized start of day timezone](/connect/customized-start-of-day#available-timezones-and-cutoffs)."}}, "description": "", "x-expandableFields": []}
```
