---
title: tunnel_config
page_id: schema-tunnel-config-98097f43
path: schemas
description: The tunnel configuration and ingress rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_config

The tunnel configuration and ingress rules.

```yaml
{"description": "The tunnel configuration and ingress rules.", "type": "object", "properties": {"ingress": {"description": "List of public hostname definitions. At least one ingress rule needs to be defined for the tunnel.", "type": "array", "items": {"$ref": "#/components/schemas/tunnel_ingressRule"}, "minItems": 1}, "originRequest": {"$ref": "#/components/schemas/tunnel_originRequest"}, "warp-routing": {"description": "Enable private network access from WARP users to private network routes. This is enabled if the tunnel has an assigned route.", "type": "object", "deprecated": true, "properties": {"enabled": {"type": "boolean"}}, "readOnly": true, "x-stainless-deprecation-message": "This field is ignored by cloudflared since version 2023.10.0.", "x-stainless-skip": true}}}
```
