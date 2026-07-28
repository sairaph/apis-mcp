---
title: tunnel_mesh_configuration_response
page_id: schema-tunnel-mesh-configuration-response-a22e674f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_mesh_configuration_response

```yaml
{"type": "object", "properties": {"config": {"description": "Provider-specific configuration. Present for `aws` and `local` modes.", "type": "object", "nullable": true, "oneOf": [{"$ref": "#/components/schemas/tunnel_mesh_aws_config"}, {"$ref": "#/components/schemas/tunnel_mesh_local_config"}]}, "configuration_version": {"description": "Monotonically increasing configuration version, incremented on each PUT.", "type": "integer"}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "ha_mode": {"$ref": "#/components/schemas/tunnel_mesh_ha_mode"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}, "updated_at": {"$ref": "#/components/schemas/tunnel_updated_at"}}, "required": ["tunnel_id", "configuration_version", "ha_mode", "created_at"]}
```
