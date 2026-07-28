---
title: dns-records_settings
page_id: schema-dns-records-settings-e7b1ce08
path: schemas
description: Settings for the DNS record.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_settings

Settings for the DNS record.

```yaml
{"description": "Settings for the DNS record.", "type": "object", "properties": {"ipv4_only": {"description": "When enabled, only A records will be generated, and AAAA records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.", "type": "boolean", "example": true, "default": false, "x-auditable": true}, "ipv6_only": {"description": "When enabled, only AAAA records will be generated, and A records will not be created. This setting is intended for exceptional cases. Note that this option only applies to proxied records and it has no effect on whether Cloudflare communicates with the origin using IPv4 or IPv6.", "type": "boolean", "example": true, "default": false, "x-auditable": true}}}
```
