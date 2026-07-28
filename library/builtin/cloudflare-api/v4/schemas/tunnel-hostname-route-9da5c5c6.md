---
title: tunnel_hostname_route
page_id: schema-tunnel-hostname-route-9da5c5c6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_hostname_route

```yaml
{"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_hostname_comment"}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "deleted_at": {"$ref": "#/components/schemas/tunnel_deleted_at"}, "hostname": {"$ref": "#/components/schemas/tunnel_hostname"}, "id": {"$ref": "#/components/schemas/tunnel_hostname_route_id"}, "tun_type": {"allOf": [{"$ref": "#/components/schemas/tunnel_tunnel_type"}, {"enum": ["cfd_tunnel", "warp_connector"], "type": "string"}]}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id-3"}, "tunnel_name": {"$ref": "#/components/schemas/tunnel_tunnel_name-2"}}}
```
