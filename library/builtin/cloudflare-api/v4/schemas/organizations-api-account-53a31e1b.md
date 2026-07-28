---
title: organizations-api_Account
page_id: schema-organizations-api-account-53a31e1b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_Account

```yaml
{"type": "object", "properties": {"created_on": {"type": "string", "format": "date-time"}, "id": {"type": "string"}, "name": {"type": "string", "nullable": true}, "settings": {"type": "object", "properties": {"abuse_contact_email": {"type": "string", "nullable": true}, "access_approval_expiry": {"type": "string", "format": "date-time", "nullable": true}, "api_access_enabled": {"type": "boolean", "nullable": true}, "default_nameservers": {"description": "Use [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings) instead. Deprecated.", "type": "string", "deprecated": true, "nullable": true}, "enforce_twofactor": {"type": "boolean", "nullable": true}, "use_account_custom_ns_by_default": {"description": "Use [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings) instead. Deprecated.", "type": "boolean", "deprecated": true, "nullable": true}}, "required": ["enforce_twofactor", "api_access_enabled", "access_approval_expiry", "abuse_contact_email", "use_account_custom_ns_by_default", "default_nameservers"]}, "type": {"type": "string", "enum": ["standard", "enterprise"]}}, "required": ["id", "name", "type", "settings", "created_on"]}
```
