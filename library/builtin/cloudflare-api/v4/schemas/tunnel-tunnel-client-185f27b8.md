---
title: tunnel_tunnel_client
page_id: schema-tunnel-tunnel-client-185f27b8
path: schemas
description: A client (typically cloudflared) that maintains connections to a Cloudflare data center.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_tunnel_client

A client (typically cloudflared) that maintains connections to a Cloudflare data center.

```yaml
{"description": "A client (typically cloudflared) that maintains connections to a Cloudflare data center.", "type": "object", "properties": {"arch": {"$ref": "#/components/schemas/tunnel_arch"}, "config_version": {"$ref": "#/components/schemas/tunnel_config_version"}, "conns": {"$ref": "#/components/schemas/tunnel_connections"}, "features": {"$ref": "#/components/schemas/tunnel_features"}, "id": {"$ref": "#/components/schemas/tunnel_connection_id"}, "run_at": {"$ref": "#/components/schemas/tunnel_run_at"}, "version": {"$ref": "#/components/schemas/tunnel_version"}}}
```
