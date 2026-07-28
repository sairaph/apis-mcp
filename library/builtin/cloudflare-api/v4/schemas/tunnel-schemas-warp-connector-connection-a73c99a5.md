---
title: tunnel_schemas-warp-connector-connection
page_id: schema-tunnel-schemas-warp-connector-connection-a73c99a5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_schemas-warp-connector-connection

```yaml
{"properties": {"client_id": {"$ref": "#/components/schemas/tunnel_client_id"}, "client_version": {"$ref": "#/components/schemas/tunnel_version"}, "colo_name": {"$ref": "#/components/schemas/tunnel_colo_name"}, "id": {"$ref": "#/components/schemas/tunnel_connection_id"}, "opened_at": {"description": "Timestamp of when the connection was established.", "type": "string", "format": "date-time", "example": "2021-01-25T18:22:34.317854Z"}, "origin_ip": {"description": "The public IP address of the host running WARP Connector.", "allOf": [{"$ref": "#/components/schemas/tunnel_ip"}]}}}
```
