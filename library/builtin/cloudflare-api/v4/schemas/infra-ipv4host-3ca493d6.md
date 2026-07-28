---
title: infra_IPv4Host
page_id: schema-infra-ipv4host-3ca493d6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_IPv4Host

```yaml
{"type": "object", "properties": {"ipv4": {"type": "string", "example": "10.0.0.1"}, "network": {"$ref": "#/components/schemas/infra_Network"}}, "example": {"ipv4": "10.0.0.1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "required": ["ipv4", "network"]}
```
