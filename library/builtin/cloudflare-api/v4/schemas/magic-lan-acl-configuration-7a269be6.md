---
title: magic_lan-acl-configuration
page_id: schema-magic-lan-acl-configuration-7a269be6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_lan-acl-configuration

```yaml
{"type": "object", "properties": {"lan_id": {"description": "The identifier for the LAN you want to create an ACL policy with.", "type": "string"}, "lan_name": {"description": "The name of the LAN based on the provided lan_id.", "type": "string"}, "port_ranges": {"description": "Array of port ranges on the provided LAN that will be included in the ACL. If no ports or port rangess are provided, communication on any port on this LAN is allowed.", "type": "array", "items": {"$ref": "#/components/schemas/magic_acl-port-range"}}, "ports": {"description": "Array of ports on the provided LAN that will be included in the ACL. If no ports or port ranges are provided, communication on any port on this LAN is allowed.", "type": "array", "items": {"$ref": "#/components/schemas/magic_port"}}, "subnets": {"description": "Array of subnet IPs within the LAN that will be included in the ACL. If no subnets are provided, communication on any subnets on this LAN are allowed.", "type": "array", "items": {"$ref": "#/components/schemas/magic_acl-subnet"}}}, "required": ["lan_id"]}
```
