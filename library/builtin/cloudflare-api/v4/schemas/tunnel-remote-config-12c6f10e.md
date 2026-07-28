---
title: tunnel_remote_config
page_id: schema-tunnel-remote-config-12c6f10e
path: schemas
description: If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tunnel_remote_config

If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.

```yaml
{"description": "If `true`, the tunnel can be configured remotely from the Zero Trust dashboard. If `false`, the tunnel must be configured locally on the origin machine.", "type": "boolean", "example": true, "deprecated": true, "x-auditable": true, "x-stainless-deprecation-message": "Use the config_src field instead.", "x-stainless-ignore": true}
```
