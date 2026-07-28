---
title: firewall_cidr_configuration
page_id: schema-firewall-cidr-configuration-8332ba1e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_cidr_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `ip_range` when specifying an IP address range in the rule.", "type": "string", "example": "ip_range", "enum": ["ip_range"]}, "value": {"description": "The IP address range to match. You can only use prefix lengths `/16` and `/24` for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.", "type": "string", "example": "198.51.100.4/16"}}, "title": "An IP address range configuration."}
```
