---
title: magic_lan
page_id: schema-magic-lan-c49915f5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_lan

```yaml
{"type": "object", "properties": {"bond_id": {"$ref": "#/components/schemas/magic_bond_id"}, "ha_link": {"description": "mark true to use this LAN for HA probing. only works for site with HA turned on. only one LAN can be set as the ha_link.", "type": "boolean"}, "id": {"$ref": "#/components/schemas/magic_identifier"}, "is_breakout": {"description": "mark true to use this LAN for source-based breakout traffic", "type": "boolean"}, "is_prioritized": {"description": "mark true to use this LAN for source-based prioritized traffic", "type": "boolean"}, "name": {"type": "string"}, "nat": {"$ref": "#/components/schemas/magic_nat"}, "physport": {"$ref": "#/components/schemas/magic_port"}, "routed_subnets": {"type": "array", "items": {"$ref": "#/components/schemas/magic_routed_subnet"}}, "site_id": {"$ref": "#/components/schemas/magic_identifier"}, "static_addressing": {"$ref": "#/components/schemas/magic_lan_static_addressing"}, "vlan_tag": {"$ref": "#/components/schemas/magic_vlan_tag"}}}
```
