---
title: magic_ipsec_tunnel_add_single_request
page_id: schema-magic-ipsec-tunnel-add-single-request-9a5ec03d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec_tunnel_add_single_request

```yaml
{"type": "object", "properties": {"automatic_return_routing": {"$ref": "#/components/schemas/magic_automatic_return_routing"}, "bgp": {"$ref": "#/components/schemas/magic_bgp_config"}, "cloudflare_endpoint": {"$ref": "#/components/schemas/magic_cloudflare_ipsec_endpoint"}, "custom_remote_identities": {"$ref": "#/components/schemas/magic_custom_remote_identities"}, "customer_endpoint": {"$ref": "#/components/schemas/magic_customer_ipsec_endpoint"}, "description": {"$ref": "#/components/schemas/magic_components-schemas-description"}, "health_check": {"$ref": "#/components/schemas/magic_tunnel_health_check"}, "interface_address": {"$ref": "#/components/schemas/magic_interface_address"}, "interface_address6": {"$ref": "#/components/schemas/magic_interface_address6"}, "name": {"$ref": "#/components/schemas/magic_ipsec_tunnel_name"}, "psk": {"$ref": "#/components/schemas/magic_psk"}, "replay_protection": {"$ref": "#/components/schemas/magic_replay_protection"}}, "required": ["name", "cloudflare_endpoint", "interface_address"]}
```
