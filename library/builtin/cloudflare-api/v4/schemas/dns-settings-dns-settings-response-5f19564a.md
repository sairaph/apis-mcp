---
title: dns-settings_dns-settings-response
page_id: schema-dns-settings-dns-settings-response-5f19564a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-settings_dns-settings-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-settings_dns-settings-base"}, {"properties": {"internal_dns": {"$ref": "#/components/schemas/dns-settings_internal_dns_response"}, "soa": {"$ref": "#/components/schemas/dns-settings_soa-response"}}, "required": ["flatten_all_cnames", "foundation_dns", "multi_provider", "secondary_overrides", "soa", "ns_ttl", "zone_mode", "internal_dns"], "type": "object"}]}
```
