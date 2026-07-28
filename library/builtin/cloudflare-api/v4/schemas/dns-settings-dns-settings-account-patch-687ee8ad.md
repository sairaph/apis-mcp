---
title: dns-settings_dns-settings-account-patch
page_id: schema-dns-settings-dns-settings-account-patch-687ee8ad
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_dns-settings-account-patch

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns-settings-patch"}, {"properties": {"nameservers": {"description": "Settings determining the nameservers through which the zone should be available.", "type": "object", "properties": {"type": {"description": "Nameserver type", "type": "string", "example": "cloudflare.standard", "enum": ["cloudflare.standard", "cloudflare.advanced", "cloudflare.standard.random", "custom.account", "custom.tenant"], "x-auditable": true}}}}, "type": "object"}]}
```
