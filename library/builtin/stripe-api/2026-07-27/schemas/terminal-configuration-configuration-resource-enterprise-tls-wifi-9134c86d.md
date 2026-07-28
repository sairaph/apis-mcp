---
title: terminal_configuration_configuration_resource_enterprise_tls_wifi
page_id: schema-terminal-configuration-configuration-resource-enterprise-tls-wifi-9134c86d
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# terminal_configuration_configuration_resource_enterprise_tls_wifi

```yaml
{"title": "TerminalConfigurationConfigurationResourceEnterpriseTLSWifi", "required": ["client_certificate_file", "private_key_file", "ssid"], "type": "object", "properties": {"ca_certificate_file": {"maxLength": 5000, "type": "string", "description": "A File ID representing a PEM file containing the server certificate"}, "client_certificate_file": {"maxLength": 5000, "type": "string", "description": "A File ID representing a PEM file containing the client certificate"}, "private_key_file": {"maxLength": 5000, "type": "string", "description": "A File ID representing a PEM file containing the client RSA private key"}, "private_key_file_password": {"maxLength": 5000, "type": "string", "description": "Password for the private key file"}, "ssid": {"maxLength": 5000, "type": "string", "description": "Name of the WiFi network"}}, "description": "", "x-expandableFields": []}
```
