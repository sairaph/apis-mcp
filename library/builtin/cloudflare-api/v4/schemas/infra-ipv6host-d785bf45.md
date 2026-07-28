---
title: infra_IPv6Host
page_id: schema-infra-ipv6host-d785bf45
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_IPv6Host

```yaml
{"type": "object", "properties": {"ipv6": {"type": "string", "example": "fe80::1"}, "network": {"$ref": "#/components/schemas/infra_Network"}}, "example": {"ipv6": "fe80::1", "network": {"tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "required": ["ipv6", "network"]}
```
