---
title: zero-trust-gateway_ipv4_networks
page_id: schema-zero-trust-gateway-ipv4-networks-085b1ab8
path: schemas
description: Specify the list of network ranges from which requests at this location originate. The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_ipv4_networks

Specify the list of network ranges from which requests at this location originate. The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.

```yaml
{"description": "Specify the list of network ranges from which requests at this location originate. The list takes effect only if it is non-empty and the IPv4 endpoint is enabled for this location.", "type": "array", "items": {"$ref": "#/components/schemas/zero-trust-gateway_ipv4_network"}, "nullable": true, "x-stainless-terraform-configurability": "optional"}
```
