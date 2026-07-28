---
title: firewall_ipv6_configuration
page_id: schema-firewall-ipv6-configuration-f1f9fa99
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_ipv6_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `ip6` when specifying an IPv6 address in the rule.", "type": "string", "example": "ip6", "enum": ["ip6"]}, "value": {"description": "The IPv6 address to match.", "type": "string", "example": "2001:DB8:100::CF"}}, "title": "An IPv6 address configuration."}
```
