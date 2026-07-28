---
title: zones_cache-rules_aegis
page_id: schema-zones-cache-rules-aegis-4ed7931e
path: schemas
description: Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your account so that you can increase your origin security by only allowing traffic from a small list of IP addresses.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache-rules_aegis

Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your account so that you can increase your origin security by only allowing traffic from a small list of IP addresses.

```yaml
{"description": "Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your account so that you can increase your origin security by only allowing traffic from a small list of IP addresses.", "type": "object", "allOf": [{"$ref": "#/components/schemas/zones_cache-rules_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "type": "string", "example": "aegis", "enum": ["aegis"], "x-auditable": true}, "value": {"$ref": "#/components/schemas/zones_cache-rules_aegis_value"}}, "type": "object"}], "title": "Aegis"}
```
