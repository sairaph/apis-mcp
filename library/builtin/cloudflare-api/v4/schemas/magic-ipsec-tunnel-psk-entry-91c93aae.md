---
title: magic_ipsec_tunnel_psk_entry
page_id: schema-magic-ipsec-tunnel-psk-entry-91c93aae
path: schemas
description: A PSK entry for a specific IPsec tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_ipsec_tunnel_psk_entry

A PSK entry for a specific IPsec tunnel.

```yaml
{"description": "A PSK entry for a specific IPsec tunnel.", "type": "object", "properties": {"id": {"description": "The ID of the IPsec tunnel.", "allOf": [{"$ref": "#/components/schemas/magic_identifier"}]}, "psk": {"$ref": "#/components/schemas/magic_psk"}}, "required": ["id", "psk"]}
```
