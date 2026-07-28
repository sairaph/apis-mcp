---
title: zero-trust-gateway_gateway-account-logging-settings
page_id: schema-zero-trust-gateway-gateway-account-logging-settings-43700749
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_gateway-account-logging-settings

```yaml
{"type": "object", "properties": {"redact_pii": {"description": "Indicate whether to redact personally identifiable information from activity logging (PII fields include source IP, user email, user ID, device ID, URL, referrer, and user agent).", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "settings_by_rule_type": {"description": "Configure logging settings for each rule type.", "type": "object", "properties": {"dns": {"description": "Configure logging settings for DNS firewall.", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_account-log-options"}]}, "http": {"description": "Configure logging settings for HTTP/HTTPS firewall.", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_account-log-options"}]}, "l4": {"description": "Configure logging settings for Network firewall.", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_account-log-options"}]}}}}}
```
