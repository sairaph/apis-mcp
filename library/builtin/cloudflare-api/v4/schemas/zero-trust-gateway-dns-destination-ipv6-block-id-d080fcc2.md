---
title: zero-trust-gateway_dns_destination_ipv6_block_id
page_id: schema-zero-trust-gateway-dns-destination-ipv6-block-id-d080fcc2
path: schemas
description: Specify the UUID of the IPv6 block brought to the gateway so that this location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6) block rather than the standard Cloudflare IPv6 block.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_dns_destination_ipv6_block_id

Specify the UUID of the IPv6 block brought to the gateway so that this location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6) block rather than the standard Cloudflare IPv6 block.

```yaml
{"description": "Specify the UUID of the IPv6 block brought to the gateway so that this location's IPv6 address is allocated from the Bring Your Own IPv6 (BYOIPv6) block rather than the standard Cloudflare IPv6 block.", "type": "string", "example": "b08f7231-d458-495c-98ef-190604c9ee83", "nullable": true, "x-auditable": true, "x-stainless-terraform-configurability": "optional"}
```
