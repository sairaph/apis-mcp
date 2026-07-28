---
title: firewall_schemas-ip_configuration
page_id: schema-firewall-schemas-ip-configuration-68954b77
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_schemas-ip_configuration

```yaml
{"type": "object", "properties": {"target": {"description": "The configuration target. You must set the target to `ip` when specifying an IP address in the Zone Lockdown rule.", "type": "string", "example": "ip", "enum": ["ip"]}, "value": {"description": "The IP address to match. This address will be compared to the IP address of incoming requests.", "type": "string", "example": "198.51.100.4"}}, "title": "An IP address configuration."}
```
