---
title: dns-settings_dns-settings-zone-response
page_id: schema-dns-settings-dns-settings-zone-response-c113721f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_dns-settings-zone-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns-settings-response"}, {"properties": {"nameservers": {"description": "Settings determining the nameservers through which the zone should be available.", "type": "object", "properties": {"ns_set": {"description": "Configured nameserver set to be used for this zone", "type": "integer", "example": 1, "maximum": 5, "minimum": 1, "x-auditable": true}, "type": {"description": "Nameserver type", "type": "string", "example": "cloudflare.standard", "enum": ["cloudflare.standard", "cloudflare.advanced", "custom.account", "custom.tenant", "custom.zone"], "x-auditable": true}}, "required": ["type"]}}, "required": ["nameservers"], "type": "object"}]}
```
