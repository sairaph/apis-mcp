---
title: zero-trust-gateway_gateway-account-settings
page_id: schema-zero-trust-gateway-gateway-account-settings-e5ffeba6
path: schemas
description: Specify account settings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_gateway-account-settings

Specify account settings.

```yaml
{"description": "Specify account settings.", "properties": {"settings": {"description": "Specify account settings.", "type": "object", "properties": {"activity_log": {"$ref": "#/components/schemas/zero-trust-gateway_activity-log-settings"}, "antivirus": {"$ref": "#/components/schemas/zero-trust-gateway_anti-virus-settings"}, "block_page": {"$ref": "#/components/schemas/zero-trust-gateway_block-page-settings"}, "body_scanning": {"$ref": "#/components/schemas/zero-trust-gateway_body-scanning-settings"}, "browser_isolation": {"$ref": "#/components/schemas/zero-trust-gateway_browser-isolation-settings"}, "certificate": {"$ref": "#/components/schemas/zero-trust-gateway_certificate-settings"}, "custom_certificate": {"$ref": "#/components/schemas/zero-trust-gateway_custom-certificate-settings"}, "extended_email_matching": {"$ref": "#/components/schemas/zero-trust-gateway_extended-email-matching"}, "fips": {"$ref": "#/components/schemas/zero-trust-gateway_fips-settings"}, "host_selector": {"$ref": "#/components/schemas/zero-trust-gateway_host-selector-settings"}, "inspection": {"$ref": "#/components/schemas/zero-trust-gateway_inspection-settings"}, "max_ttl_secs": {"$ref": "#/components/schemas/zero-trust-gateway_max-ttl-secs"}, "protocol_detection": {"$ref": "#/components/schemas/zero-trust-gateway_protocol-detection"}, "sandbox": {"$ref": "#/components/schemas/zero-trust-gateway_sandbox"}, "tls_decrypt": {"$ref": "#/components/schemas/zero-trust-gateway_tls-settings"}}}}}
```
