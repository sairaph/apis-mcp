---
title: firewall_schemas-cidr_configuration
page_id: schema-firewall-schemas-cidr-configuration-8bcae809
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_schemas-cidr_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `ip_range` when specifying an IP address range in the Zone Lockdown rule.", "type": "string", "example": "ip_range", "enum": ["ip_range"]}, "value": {"description": "The IP address range to match. You can only use prefix lengths `/16` and `/24`.", "type": "string", "example": "198.51.100.4/16"}}, "title": "An IP address range configuration."}
```
