---
title: tunnel_ingressRule
page_id: schema-tunnel-ingressrule-77c7e2ae
path: schemas
description: Public hostname
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_ingressRule

Public hostname

```yaml
{"description": "Public hostname", "type": "object", "properties": {"hostname": {"description": "Public hostname for this service.", "type": "string", "example": "tunnel.example.com"}, "originRequest": {"$ref": "#/components/schemas/tunnel_originRequest"}, "path": {"description": "Requests with this path route to this public hostname.", "type": "string", "example": "subpath"}, "service": {"description": "Protocol and address of destination server. Supported protocols: http://, https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively can return a HTTP status code http_status:[code] e.g. 'http_status:404'.\n", "type": "string", "example": "https://localhost:8001"}}, "required": ["hostname", "service"]}
```
