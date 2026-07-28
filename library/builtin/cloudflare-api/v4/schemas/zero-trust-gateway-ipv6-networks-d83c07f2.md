---
title: zero-trust-gateway_ipv6_networks
page_id: schema-zero-trust-gateway-ipv6-networks-d83c07f2
path: schemas
description: Specify the list of allowed source IPv6 network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_ipv6_networks

Specify the list of allowed source IPv6 network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.

```yaml
{"description": "Specify the list of allowed source IPv6 network ranges for this endpoint. When the list is empty, the endpoint allows all source IPs. The list takes effect only if the endpoint is enabled for this location.", "type": "array", "items": {"$ref": "#/components/schemas/zero-trust-gateway_ipv6_network"}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
