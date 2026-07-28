---
title: terminal_configuration_configuration_resource_device_type_specific_config
page_id: schema-terminal-configuration-configuration-resource-device-type-specific-confi-7e56aa9b
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_configuration_configuration_resource_device_type_specific_config

```yaml
{"title": "TerminalConfigurationConfigurationResourceDeviceTypeSpecificConfig", "type": "object", "properties": {"splashscreen": {"description": "A File ID representing an image to display on the reader", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/file"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/file"}]}}}, "description": "", "x-expandableFields": ["splashscreen"]}
```
