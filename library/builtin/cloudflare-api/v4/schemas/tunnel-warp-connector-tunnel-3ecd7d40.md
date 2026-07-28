---
title: tunnel_warp_connector_tunnel
page_id: schema-tunnel-warp-connector-tunnel-3ecd7d40
path: schemas
description: A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_warp_connector_tunnel

A Warp Connector Tunnel that connects your origin to Cloudflare's edge.

```yaml
{"description": "A Warp Connector Tunnel that connects your origin to Cloudflare's edge.", "type": "object", "properties": {"account_tag": {"$ref": "#/components/schemas/tunnel_account_id"}, "connections": {"$ref": "#/components/schemas/tunnel_connections_deprecated"}, "conns_active_at": {"$ref": "#/components/schemas/tunnel_conns_active_at"}, "conns_inactive_at": {"$ref": "#/components/schemas/tunnel_conns_inactive_at"}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "deleted_at": {"$ref": "#/components/schemas/tunnel_deleted_at"}, "id": {"$ref": "#/components/schemas/tunnel_tunnel_id"}, "metadata": {"$ref": "#/components/schemas/tunnel_metadata"}, "name": {"$ref": "#/components/schemas/tunnel_tunnel_name"}, "status": {"$ref": "#/components/schemas/tunnel_status"}, "tun_type": {"$ref": "#/components/schemas/tunnel_tunnel_type"}}}
```
