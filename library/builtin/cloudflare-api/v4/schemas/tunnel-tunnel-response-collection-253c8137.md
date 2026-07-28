---
title: tunnel_tunnel-response-collection
page_id: schema-tunnel-tunnel-response-collection-253c8137
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_tunnel-response-collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/tunnel_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"anyOf": [{"$ref": "#/components/schemas/tunnel_cfd_tunnel"}, {"$ref": "#/components/schemas/tunnel_warp_connector_tunnel"}]}}}, "type": "object"}]}
```
