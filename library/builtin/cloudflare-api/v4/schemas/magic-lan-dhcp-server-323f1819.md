---
title: magic_lan_dhcp_server
page_id: schema-magic-lan-dhcp-server-323f1819
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_lan_dhcp_server

```yaml
{"type": "object", "properties": {"dhcp_options": {"description": "Optional list of custom DHCP options to include in DHCP responses. Only valid when DHCP server is enabled.", "type": "array", "items": {"$ref": "#/components/schemas/magic_dhcp_option"}}, "dhcp_pool_end": {"$ref": "#/components/schemas/magic_ip-address"}, "dhcp_pool_start": {"$ref": "#/components/schemas/magic_ip-address"}, "dns_server": {"$ref": "#/components/schemas/magic_ip-address"}, "dns_servers": {"type": "array", "items": {"$ref": "#/components/schemas/magic_ip-address"}}, "reservations": {"description": "Mapping of MAC addresses to IP addresses", "type": "object", "example": {"00:11:22:33:44:55": "192.0.2.100", "AA:BB:CC:DD:EE:FF": "192.168.1.101"}, "additionalProperties": {"description": "IP address associated with the MAC address", "type": "string"}}}}
```
