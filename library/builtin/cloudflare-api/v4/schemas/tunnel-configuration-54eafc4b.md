---
title: tunnel_configuration
page_id: schema-tunnel-configuration-54eafc4b
path: schemas
description: Cloudflare Tunnel configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_configuration

Cloudflare Tunnel configuration

```yaml
{"description": "Cloudflare Tunnel configuration", "type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/tunnel_identifier"}, "config": {"$ref": "#/components/schemas/tunnel_config"}, "created_at": {"$ref": "#/components/schemas/tunnel_timestamp"}, "source": {"$ref": "#/components/schemas/tunnel_config_src-2"}, "tunnel_id": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}, "version": {"$ref": "#/components/schemas/tunnel_config_version-2"}}}
```
