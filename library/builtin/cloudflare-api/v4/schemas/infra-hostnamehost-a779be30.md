---
title: infra_HostnameHost
page_id: schema-infra-hostnamehost-a779be30
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# infra_HostnameHost

```yaml
{"type": "object", "properties": {"hostname": {"type": "string", "example": "example.com"}, "resolver_network": {"$ref": "#/components/schemas/infra_ResolverNetwork"}}, "example": {"hostname": "example.com", "resolver_network": {"resolver_ips": ["10.0.0.1"], "tunnel_id": "0191dce4-9ab4-7fce-b660-8e5dec5172da"}}, "required": ["hostname", "resolver_network"]}
```
