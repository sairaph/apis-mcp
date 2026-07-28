---
title: zero-trust-gateway_dns_resolver_settings_v6
page_id: schema-zero-trust-gateway-dns-resolver-settings-v6-1d6571b5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_dns_resolver_settings_v6

```yaml
{"type": "object", "properties": {"ip": {"description": "Specify the IPv6 address of the upstream resolver.", "type": "string", "example": "2001:DB8::", "x-auditable": true}, "port": {"description": "Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.", "type": "integer", "example": 5053, "x-auditable": true}, "route_through_private_network": {"description": "Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.", "type": "boolean", "example": true, "x-auditable": true}, "vnet_id": {"description": "Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.", "type": "string", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415", "x-auditable": true}}, "required": ["ip"], "x-stainless-terraform-configurability": "computed_optional"}
```
