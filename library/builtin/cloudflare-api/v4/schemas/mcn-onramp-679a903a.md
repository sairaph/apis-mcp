---
title: mcn_onramp
page_id: schema-mcn-onramp-679a903a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_onramp

```yaml
{"type": "object", "properties": {"attached_hubs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "attached_vpcs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "cloud_asn": {"type": "integer", "format": "uint32", "x-auditable": true}, "cloud_type": {"$ref": "#/components/schemas/mcn_onramp_cloud_type"}, "description": {"type": "string", "x-auditable": true}, "dynamic_routing": {"type": "boolean", "x-auditable": true}, "hub": {"$ref": "#/components/schemas/mcn_resource_id"}, "id": {"$ref": "#/components/schemas/mcn_onramp_id"}, "install_routes_in_cloud": {"type": "boolean", "x-auditable": true}, "install_routes_in_magic_wan": {"type": "boolean", "x-auditable": true}, "last_applied_at": {"type": "string", "x-auditable": true}, "last_exported_at": {"type": "string", "x-auditable": true}, "last_planned_at": {"type": "string", "x-auditable": true}, "manage_hub_to_hub_attachments": {"type": "boolean", "x-auditable": true}, "manage_vpc_to_hub_attachments": {"type": "boolean", "x-auditable": true}, "name": {"type": "string", "x-auditable": true}, "planned_monthly_cost_estimate": {"$ref": "#/components/schemas/mcn_cost_diff"}, "planned_resources": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_diff"}}, "planned_resources_unavailable": {"type": "boolean", "x-auditable": true}, "post_apply_monthly_cost_estimate": {"$ref": "#/components/schemas/mcn_cost"}, "post_apply_resources": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/mcn_resource_details"}}, "post_apply_resources_unavailable": {"type": "boolean", "x-auditable": true}, "region": {"type": "string", "x-auditable": true}, "status": {"$ref": "#/components/schemas/mcn_onramp_status"}, "type": {"$ref": "#/components/schemas/mcn_onramp_type"}, "updated_at": {"type": "string"}, "vpc": {"$ref": "#/components/schemas/mcn_resource_id"}, "vpcs_by_id": {"type": "object", "additionalProperties": {"$ref": "#/components/schemas/mcn_resource_details"}}, "vpcs_by_id_unavailable": {"description": "The list of vpc IDs for which resource details failed to generate.", "type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}}, "required": ["id", "name", "type", "cloud_type", "install_routes_in_cloud", "install_routes_in_magic_wan", "dynamic_routing", "updated_at"]}
```
