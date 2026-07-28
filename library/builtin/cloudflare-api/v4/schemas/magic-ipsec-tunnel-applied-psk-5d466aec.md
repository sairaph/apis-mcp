---
title: magic_ipsec_tunnel_applied_psk
page_id: schema-magic-ipsec-tunnel-applied-psk-5d466aec
path: schemas
description: A successfully applied PSK for an IPsec tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec_tunnel_applied_psk

A successfully applied PSK for an IPsec tunnel.

```yaml
{"description": "A successfully applied PSK for an IPsec tunnel.", "type": "object", "properties": {"ipsec_id": {"description": "The IKE identifier used for this tunnel on the Cloudflare edge.", "type": "string", "example": "12345_abc123def4567890abcdef1234567890"}, "ipsec_tunnel_id": {"$ref": "#/components/schemas/magic_identifier"}, "psk": {"$ref": "#/components/schemas/magic_psk"}, "psk_metadata": {"$ref": "#/components/schemas/magic_psk_metadata"}}, "required": ["ipsec_tunnel_id", "ipsec_id", "psk", "psk_metadata"]}
```
