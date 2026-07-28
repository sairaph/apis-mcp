---
title: mcn_create_onramp_request
page_id: schema-mcn-create-onramp-request-2372fa23
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_create_onramp_request

```yaml
{"type": "object", "properties": {"adopted_hub_id": {"$ref": "#/components/schemas/mcn_resource_id"}, "attached_hubs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "attached_vpcs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "cloud_asn": {"description": "Sets the cloud-side ASN. If unset or zero, the cloud's default ASN takes effect.", "type": "integer", "format": "uint32", "x-auditable": true}, "cloud_type": {"$ref": "#/components/schemas/mcn_onramp_cloud_type"}, "description": {"type": "string"}, "dynamic_routing": {"description": "Enables BGP routing. When enabling this feature, set both install_routes_in_cloud and install_routes_in_magic_wan to false.", "type": "boolean", "x-auditable": true}, "hub_provider_id": {"$ref": "#/components/schemas/mcn_provider_id"}, "install_routes_in_cloud": {"type": "boolean", "x-auditable": true}, "install_routes_in_magic_wan": {"type": "boolean", "x-auditable": true}, "manage_hub_to_hub_attachments": {"type": "boolean", "x-auditable": true}, "manage_vpc_to_hub_attachments": {"type": "boolean", "x-auditable": true}, "name": {"type": "string"}, "region": {"type": "string", "x-auditable": true}, "type": {"$ref": "#/components/schemas/mcn_onramp_type"}, "vpc": {"$ref": "#/components/schemas/mcn_resource_id"}}, "required": ["name", "type", "cloud_type", "install_routes_in_cloud", "install_routes_in_magic_wan", "dynamic_routing"]}
```
