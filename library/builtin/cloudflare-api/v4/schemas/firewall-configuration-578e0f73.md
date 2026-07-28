---
title: firewall_configuration
page_id: schema-firewall-configuration-578e0f73
path: schemas
description: The rule configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_configuration

The rule configuration.

```yaml
{"description": "The rule configuration.", "type": "object", "oneOf": [{"$ref": "#/components/schemas/firewall_ip_configuration"}, {"$ref": "#/components/schemas/firewall_ipv6_configuration"}, {"$ref": "#/components/schemas/firewall_cidr_configuration"}, {"$ref": "#/components/schemas/firewall_asn_configuration"}, {"$ref": "#/components/schemas/firewall_country_configuration"}]}
```
