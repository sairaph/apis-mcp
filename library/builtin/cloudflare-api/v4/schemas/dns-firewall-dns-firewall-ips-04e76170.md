---
title: dns-firewall_dns_firewall_ips
page_id: schema-dns-firewall-dns-firewall-ips-04e76170
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_dns_firewall_ips

```yaml
{"type": "array", "items": {"anyOf": [{"description": "Cloudflare-assigned DNS IPv4 address", "example": "203.0.113.1", "format": "ipv4", "type": "string", "x-auditable": true}, {"description": "Cloudflare-assigned DNS IPv6 address", "example": "2001:DB8:ab::CF", "format": "ipv6", "type": "string", "x-auditable": true}], "type": "string"}, "example": ["203.0.113.1", "203.0.113.254", "2001:DB8:AB::CF", "2001:DB8:CD::CF"], "x-stainless-collection-type": "set"}
```
