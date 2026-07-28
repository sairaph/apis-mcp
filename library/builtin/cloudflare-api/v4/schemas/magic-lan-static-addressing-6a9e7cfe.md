---
title: magic_lan_static_addressing
page_id: schema-magic-lan-static-addressing-6a9e7cfe
path: schemas
description: If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_lan_static_addressing

If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.

```yaml
{"description": "If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.", "type": "object", "properties": {"address": {"$ref": "#/components/schemas/magic_cidr"}, "dhcp_relay": {"$ref": "#/components/schemas/magic_lan_dhcp_relay"}, "dhcp_server": {"$ref": "#/components/schemas/magic_lan_dhcp_server"}, "secondary_address": {"$ref": "#/components/schemas/magic_cidr"}, "virtual_address": {"$ref": "#/components/schemas/magic_cidr"}}, "required": ["address"]}
```
