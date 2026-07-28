---
title: tunnel_tunnel_warp_connector_client
page_id: schema-tunnel-tunnel-warp-connector-client-eb84025a
path: schemas
description: A WARP Connector client that maintains a connection to a Cloudflare data center.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_tunnel_warp_connector_client

A WARP Connector client that maintains a connection to a Cloudflare data center.

```yaml
{"description": "A WARP Connector client that maintains a connection to a Cloudflare data center.", "type": "object", "properties": {"arch": {"$ref": "#/components/schemas/tunnel_arch"}, "conns": {"$ref": "#/components/schemas/tunnel_warp_connector_connections"}, "features": {"$ref": "#/components/schemas/tunnel_features"}, "ha_status": {"$ref": "#/components/schemas/tunnel_ha_status"}, "id": {"$ref": "#/components/schemas/tunnel_client_id"}, "run_at": {"$ref": "#/components/schemas/tunnel_run_at"}, "version": {"$ref": "#/components/schemas/tunnel_version"}}}
```
