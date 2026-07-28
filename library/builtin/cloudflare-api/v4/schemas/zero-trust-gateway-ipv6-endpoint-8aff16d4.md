---
title: zero-trust-gateway_ipv6_endpoint
page_id: schema-zero-trust-gateway-ipv6-endpoint-8aff16d4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_ipv6_endpoint

```yaml
{"type": "object", "properties": {"enabled": {"description": "Indicate whether the IPV6 endpoint is enabled for this location.", "type": "boolean", "example": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}, "networks": {"$ref": "#/components/schemas/zero-trust-gateway_ipv6_networks"}}}
```
