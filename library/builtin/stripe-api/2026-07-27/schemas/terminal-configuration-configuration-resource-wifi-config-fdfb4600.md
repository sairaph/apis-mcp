---
title: terminal_configuration_configuration_resource_wifi_config
page_id: schema-terminal-configuration-configuration-resource-wifi-config-fdfb4600
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_configuration_configuration_resource_wifi_config

```yaml
{"title": "TerminalConfigurationConfigurationResourceWifiConfig", "required": ["type"], "type": "object", "properties": {"enterprise_eap_peap": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_enterprise_peap_wifi"}, "enterprise_eap_tls": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_enterprise_tls_wifi"}, "personal_psk": {"$ref": "#/components/schemas/terminal_configuration_configuration_resource_personal_psk_wifi"}, "type": {"type": "string", "description": "Security type of the WiFi network. The hash with the corresponding name contains the credentials for this security type.", "enum": ["enterprise_eap_peap", "enterprise_eap_tls", "personal_psk"]}}, "description": "", "x-expandableFields": ["enterprise_eap_peap", "enterprise_eap_tls", "personal_psk"]}
```
