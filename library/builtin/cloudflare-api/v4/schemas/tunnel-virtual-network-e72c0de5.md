---
title: tunnel_virtual-network
page_id: schema-tunnel-virtual-network-e72c0de5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_virtual-network

```yaml
{"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_virtual_network_comment"}, "created_at": {"$ref": "#/components/schemas/tunnel_created_at"}, "deleted_at": {"$ref": "#/components/schemas/tunnel_deleted_at"}, "id": {"$ref": "#/components/schemas/tunnel_virtual_network_id"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_is_default_network"}, "name": {"$ref": "#/components/schemas/tunnel_virtual_network_name"}}, "required": ["id", "name", "is_default_network", "comment", "created_at"]}
```
