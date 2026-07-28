---
title: tls-certificates-and-hostnames_keyless_tunnel
page_id: schema-tls-certificates-and-hostnames-keyless-tunnel-14e1affb
path: schemas
description: Configuration for using Keyless SSL through a Cloudflare Tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_keyless_tunnel

Configuration for using Keyless SSL through a Cloudflare Tunnel.

```yaml
{"description": "Configuration for using Keyless SSL through a Cloudflare Tunnel.", "type": "object", "properties": {"private_ip": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_private_ip"}, "vnet_id": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_keyless_vnet_id"}}, "required": ["private_ip", "vnet_id"]}
```
