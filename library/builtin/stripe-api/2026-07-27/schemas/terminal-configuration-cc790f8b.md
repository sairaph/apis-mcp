---
title: terminal.configuration
page_id: schema-terminal-configuration-cc790f8b
path: schemas
description: |-
    A Configurations object represents how features should be configured for terminal readers.
    For information about how to use it, see the [Terminal configurations documentation](https://docs.stripe.com/terminal/fleet/configurations-overview).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal.configuration

A Configurations object represents how features should be configured for terminal readers.
For information about how to use it, see the [Terminal configurations documentation](https://docs.stripe.com/terminal/fleet/configurations-overview).

```yaml
{"title": "TerminalConfigurationConfiguration", "required": ["id", "livemode", "object"], "type": "object", "properties": {"bbpos_wisepad3": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "bbpos_wisepos_e": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "cellular": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_cellular_config"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "is_account_default": {"type": "boolean", "description": "Whether this Configuration is the default for your account", "nullable": true}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "String indicating the name of the Configuration object, set by the user", "nullable": true}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["terminal.configuration"]}, "offline": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_offline_config"}, "reboot_window": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_reboot_window"}, "stripe_s700": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "stripe_s710": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "tipping": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_tipping"}, "verifone_m425": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "verifone_p400": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "verifone_p630": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "verifone_ux700": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "verifone_v660p": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_device_type_specific_config"}, "wifi": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_wifi_config"}}, "description": "A Configurations object represents how features should be configured for terminal readers.\nFor information about how to use it, see the [Terminal configurations documentation](https://docs.stripe.com/terminal/fleet/configurations-overview).", "x-expandableFields": ["bbpos_wisepad3", "bbpos_wisepos_e", "cellular", "offline", "reboot_window", "stripe_s700", "stripe_s710", "tipping", "verifone_m425", "verifone_p400", "verifone_p630", "verifone_ux700", "verifone_v660p", "wifi"], "x-resourceId": "terminal.configuration"}
```
