---
title: tunnel_connections_deprecated
page_id: schema-tunnel-connections-deprecated-e2b60014
path: schemas
description: The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_connections_deprecated

The Cloudflare Tunnel connections between your origin and Cloudflare's edge.

```yaml
{"description": "The Cloudflare Tunnel connections between your origin and Cloudflare's edge.", "type": "array", "items": {"$ref": "#/components/schemas/tunnel_schemas-connection"}, "deprecated": true, "x-stainless-deprecation-message": "This field will start returning an empty array. To fetch the connections of a given tunnel, please use the dedicated endpoint `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`"}
```
