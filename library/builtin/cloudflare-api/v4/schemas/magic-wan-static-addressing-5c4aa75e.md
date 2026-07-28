---
title: magic_wan_static_addressing
page_id: schema-magic-wan-static-addressing-5c4aa75e
path: schemas
description: (optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_wan_static_addressing

(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.

```yaml
{"description": "(optional) if omitted, use DHCP. Submit secondary_address when site is in high availability mode.", "type": "object", "properties": {"address": {"$ref": "#/components/schemas/magic_cidr"}, "gateway_address": {"$ref": "#/components/schemas/magic_ip-address"}, "secondary_address": {"$ref": "#/components/schemas/magic_cidr"}}, "required": ["address", "gateway_address"]}
```
