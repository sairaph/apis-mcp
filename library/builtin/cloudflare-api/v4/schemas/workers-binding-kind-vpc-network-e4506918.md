---
title: workers_binding_kind_vpc_network
page_id: schema-workers-binding-kind-vpc-network-e4506918
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_binding_kind_vpc_network

```yaml
{"type": "object", "properties": {"name": {"$ref": "#/components/schemas/workers_binding_name"}, "network_id": {"description": "Identifier of the network to bind to. Only \"cf1:network\" is currently supported. Mutually exclusive with tunnel_id.\n", "type": "string", "example": "cf1:network", "x-auditable": true}, "tunnel_id": {"description": "UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.\n", "type": "string", "example": "abcd1234-5678-90ef-ghij-klmnopqrstuv", "x-auditable": true}, "type": {"description": "The kind of resource that the binding provides.", "type": "string", "enum": ["vpc_network"], "x-auditable": true}}, "required": ["name", "type"]}
```
