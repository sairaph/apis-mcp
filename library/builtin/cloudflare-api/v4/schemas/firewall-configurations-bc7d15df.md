---
title: firewall_configurations
page_id: schema-firewall-configurations-bc7d15df
path: schemas
description: A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# firewall_configurations

A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations.

```yaml
{"description": "A list of IP addresses or CIDR ranges that will be allowed to access the URLs specified in the Zone Lockdown rule. You can include any number of `ip` or `ip_range` configurations.", "type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/firewall_schemas-ip_configuration"}, {"$ref": "#/components/schemas/firewall_schemas-cidr_configuration"}]}}
```
