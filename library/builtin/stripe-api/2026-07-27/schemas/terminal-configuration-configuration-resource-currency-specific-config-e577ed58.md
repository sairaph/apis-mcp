---
title: terminal_configuration_configuration_resource_currency_specific_config
page_id: schema-terminal-configuration-configuration-resource-currency-specific-config-e577ed58
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_configuration_configuration_resource_currency_specific_config

```yaml
{"title": "TerminalConfigurationConfigurationResourceCurrencySpecificConfig", "type": "object", "properties": {"fixed_amounts": {"type": "array", "description": "Fixed amounts displayed when collecting a tip", "nullable": true, "items": {"type": "integer"}}, "percentages": {"type": "array", "description": "Percentages displayed when collecting a tip", "nullable": true, "items": {"type": "integer"}}, "smart_tip_threshold": {"type": "integer", "description": "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed"}}, "description": "", "x-expandableFields": []}
```
