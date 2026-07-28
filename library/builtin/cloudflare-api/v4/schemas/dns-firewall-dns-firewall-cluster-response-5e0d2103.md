---
title: dns-firewall_dns-firewall-cluster-response
page_id: schema-dns-firewall-dns-firewall-cluster-response-5e0d2103
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_dns-firewall-cluster-response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-firewall_dns-firewall-cluster"}, {"properties": {"dns_firewall_ips": {"$ref": "#/components/schemas/dns-firewall_dns_firewall_ips"}, "id": {"$ref": "#/components/schemas/dns-firewall_identifier"}, "modified_on": {"$ref": "#/components/schemas/dns-firewall_modified_on"}}, "required": ["id", "dns_firewall_ips", "modified_on"], "type": "object"}], "required": ["name", "upstream_ips", "minimum_cache_ttl", "maximum_cache_ttl", "negative_cache_ttl", "deprecate_any_requests", "ecs_fallback", "ratelimit", "retries"]}
```
