---
title: magic_wan
page_id: schema-magic-wan-38c241a7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_wan

```yaml
{"type": "object", "properties": {"health_check_rate": {"description": "Magic WAN health check rate for tunnels created on this link. The default value is `mid`.", "type": "string", "example": "low", "default": "mid", "enum": ["low", "mid", "high"]}, "id": {"$ref": "#/components/schemas/magic_identifier"}, "name": {"type": "string"}, "physport": {"$ref": "#/components/schemas/magic_port"}, "priority": {"description": "Priority of WAN for traffic loadbalancing.", "type": "integer"}, "site_id": {"$ref": "#/components/schemas/magic_identifier"}, "static_addressing": {"$ref": "#/components/schemas/magic_wan_static_addressing"}, "vlan_tag": {"$ref": "#/components/schemas/magic_vlan_tag"}}}
```
