---
title: tunnel_argo-tunnel
page_id: schema-tunnel-argo-tunnel-ee9b58b8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_argo-tunnel

```yaml
{"type": "object", "properties": {"connections": {"description": "The tunnel connections between your origin and Cloudflare's edge.", "type": "array", "items": {"$ref": "#/components/schemas/tunnel_connection"}}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "deleted_at": {"$ref": "#/components/schemas/tunnel_deleted_at"}, "id": {"$ref": "#/components/schemas/tunnel_tunnel_id"}, "name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}}, "required": ["id", "name", "created_at", "connections"]}
```
