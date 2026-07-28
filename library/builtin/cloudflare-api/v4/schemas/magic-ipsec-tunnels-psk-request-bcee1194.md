---
title: magic_ipsec_tunnels_psk_request
page_id: schema-magic-ipsec-tunnels-psk-request-bcee1194
path: schemas
description: Request body for setting PSKs for multiple IPsec tunnels.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec_tunnels_psk_request

Request body for setting PSKs for multiple IPsec tunnels.

```yaml
{"description": "Request body for setting PSKs for multiple IPsec tunnels.", "type": "object", "properties": {"psks": {"description": "List of tunnel ID and PSK pairs.", "type": "array", "items": {"$ref": "#/components/schemas/magic_ipsec_tunnel_psk_entry"}}}, "required": ["psks"]}
```
