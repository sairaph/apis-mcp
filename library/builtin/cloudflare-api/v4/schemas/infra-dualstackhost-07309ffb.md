---
title: infra_DualStackHost
page_id: schema-infra-dualstackhost-07309ffb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_DualStackHost

```yaml
{"type": "object", "properties": {"ipv4": {"type": "string", "example": "10.0.0.1"}, "ipv6": {"type": "string", "example": "fe80::1"}, "network": {"$ref": "#/components/schemas/infra_Network"}}, "example": {"ipv4": "10.0.0.1", "ipv6": "fe80::1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "required": ["ipv4", "ipv6", "network"]}
```
