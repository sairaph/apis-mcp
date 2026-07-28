---
title: dns-firewall_dns-firewall-cluster-post
page_id: schema-dns-firewall-dns-firewall-cluster-post-f7248809
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_dns-firewall-cluster-post

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns-firewall-cluster"}, {"properties": {"dns_firewall_ip_count": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_ip_count"}}, "type": "object"}], "required": ["name", "upstream_ips"]}
```
