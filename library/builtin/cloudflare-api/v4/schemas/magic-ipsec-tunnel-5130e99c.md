---
title: magic_ipsec-tunnel
page_id: schema-magic-ipsec-tunnel-5130e99c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec-tunnel

```yaml
{"type": "object", "properties": {"allow_null_cipher": {"$ref": "#/components/schemas/magic_allow_null_cipher"}, "automatic_return_routing": {"$ref": "#/components/schemas/magic_automatic_return_routing"}, "bgp": {"$ref": "#/components/schemas/magic_bgp_config"}, "bgp_status": {"$ref": "#/components/schemas/magic_bgp_status_with_state"}, "cloudflare_endpoint": {"$ref": "#/components/schemas/magic_cloudflare_ipsec_endpoint"}, "created_on": {"$ref": "#/components/schemas/magic_schemas-created_on"}, "custom_remote_identities": {"$ref": "#/components/schemas/magic_custom_remote_identities"}, "customer_endpoint": {"$ref": "#/components/schemas/magic_customer_ipsec_endpoint"}, "description": {"$ref": "#/components/schemas/magic_components-schemas-description"}, "health_check": {"$ref": "#/components/schemas/magic_tunnel_health_check"}, "id": {"$ref": "#/components/schemas/magic_schemas-identifier"}, "interface_address": {"$ref": "#/components/schemas/magic_interface_address"}, "interface_address6": {"$ref": "#/components/schemas/magic_interface_address6"}, "modified_on": {"$ref": "#/components/schemas/magic_schemas-modified_on"}, "name": {"$ref": "#/components/schemas/magic_ipsec_tunnel_name"}, "psk_metadata": {"$ref": "#/components/schemas/magic_psk_metadata"}, "replay_protection": {"$ref": "#/components/schemas/magic_replay_protection"}}, "required": ["id", "name", "cloudflare_endpoint", "interface_address"]}
```
