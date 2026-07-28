---
title: dns-firewall_upstream_ips
page_id: schema-dns-firewall-upstream-ips-c0d75222
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_upstream_ips

```yaml
{"type": "array", "items": {"anyOf": [{"description": "Upstream DNS Server IPv4 address", "example": "192.0.2.1", "format": "ipv4", "type": "string", "x-auditable": true}, {"description": "Upstream DNS Server IPv6 address", "example": "2001:DB8:100::CF", "format": "ipv6", "type": "string", "x-auditable": true}], "type": "string"}, "example": ["192.0.2.1", "198.51.100.1", "2001:DB8:100::CF"], "minLength": 1, "x-stainless-collection-type": "set"}
```
