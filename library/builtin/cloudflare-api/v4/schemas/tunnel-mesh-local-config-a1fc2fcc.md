---
title: tunnel_mesh_local_config
page_id: schema-tunnel-mesh-local-config-a1fc2fcc
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_mesh_local_config

```yaml
{"type": "object", "properties": {"vips": {"description": "VIPs to assign on the CloudflareWARP interface.", "type": "array", "items": {"$ref": "#/components/schemas/tunnel_mesh_vip_entry"}, "maxItems": 8, "minItems": 1}, "vips_previous": {"description": "VIPs to clean up on demotion or version drift.", "type": "array", "items": {"$ref": "#/components/schemas/tunnel_mesh_vip_entry"}, "maxItems": 8}}, "required": ["vips"]}
```
