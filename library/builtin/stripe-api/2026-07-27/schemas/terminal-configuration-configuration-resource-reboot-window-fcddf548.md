---
title: terminal_configuration_configuration_resource_reboot_window
page_id: schema-terminal-configuration-configuration-resource-reboot-window-fcddf548
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_configuration_configuration_resource_reboot_window

```yaml
{"title": "TerminalConfigurationConfigurationResourceRebootWindow", "required": ["end_hour", "start_hour"], "type": "object", "properties": {"end_hour": {"type": "integer", "description": "Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour."}, "start_hour": {"type": "integer", "description": "Integer between 0 to 23 that represents the start hour of the reboot time window."}}, "description": "", "x-expandableFields": []}
```
