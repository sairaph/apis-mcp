---
title: tunnel_schemas-connection
page_id: schema-tunnel-schemas-connection-390703f4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_schemas-connection

```yaml
{"properties": {"client_id": {"$ref": "#/components/schemas/tunnel_client_id"}, "client_version": {"$ref": "#/components/schemas/tunnel_version"}, "colo_name": {"$ref": "#/components/schemas/tunnel_colo_name"}, "id": {"$ref": "#/components/schemas/tunnel_connection_id"}, "is_pending_reconnect": {"$ref": "#/components/schemas/tunnel_is_pending_reconnect"}, "opened_at": {"description": "Timestamp of when the connection was established.", "type": "string", "format": "date-time", "example": "2021-01-25T18:22:34.317854Z"}, "origin_ip": {"description": "The public IP address of the host running cloudflared.", "allOf": [{"$ref": "#/components/schemas/tunnel_ip"}]}, "uuid": {"$ref": "#/components/schemas/tunnel_connection_id"}}}
```
