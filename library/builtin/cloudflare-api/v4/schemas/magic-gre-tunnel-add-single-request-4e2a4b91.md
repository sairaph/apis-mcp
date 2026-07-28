---
title: magic_gre_tunnel_add_single_request
page_id: schema-magic-gre-tunnel-add-single-request-4e2a4b91
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_gre_tunnel_add_single_request

```yaml
{"type": "object", "properties": {"automatic_return_routing": {"$ref": "#/components/schemas/magic_automatic_return_routing"}, "cloudflare_gre_endpoint": {"$ref": "#/components/schemas/magic_cloudflare_gre_endpoint"}, "customer_gre_endpoint": {"$ref": "#/components/schemas/magic_customer_gre_endpoint"}, "description": {"$ref": "#/components/schemas/magic_schemas-description"}, "health_check": {"$ref": "#/components/schemas/magic_tunnel_health_check"}, "interface_address": {"$ref": "#/components/schemas/magic_interface_address"}, "interface_address6": {"$ref": "#/components/schemas/magic_interface_address6"}, "mtu": {"$ref": "#/components/schemas/magic_mtu"}, "name": {"$ref": "#/components/schemas/magic_gre_tunnel_name"}, "ttl": {"$ref": "#/components/schemas/magic_ttl"}}, "required": ["name", "customer_gre_endpoint", "cloudflare_gre_endpoint", "interface_address"]}
```
