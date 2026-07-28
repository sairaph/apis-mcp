---
title: mcn_update_onramp_request
page_id: schema-mcn-update-onramp-request-40ac9447
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mcn_update_onramp_request

```yaml
{"type": "object", "properties": {"attached_hubs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "attached_vpcs": {"type": "array", "items": {"$ref": "#/components/schemas/mcn_resource_id"}}, "description": {"type": "string"}, "install_routes_in_cloud": {"type": "boolean", "x-auditable": true}, "install_routes_in_magic_wan": {"type": "boolean", "x-auditable": true}, "manage_hub_to_hub_attachments": {"type": "boolean"}, "manage_vpc_to_hub_attachments": {"type": "boolean"}, "name": {"type": "string"}, "vpc": {"$ref": "#/components/schemas/mcn_resource_id"}}}
```
