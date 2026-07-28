---
title: smartshield_tcp_config
page_id: schema-smartshield-tcp-config-991c455f
path: schemas
description: Parameters specific to TCP health check.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# smartshield_tcp_config

Parameters specific to TCP health check.

```yaml
{"description": "Parameters specific to TCP health check.", "type": "object", "properties": {"method": {"description": "The TCP connection method to use for the health check.", "type": "string", "default": "connection_established", "enum": ["connection_established"], "x-auditable": true}, "port": {"description": "Port number to connect to for the health check. Defaults to 80.", "type": "integer", "default": 80, "x-auditable": true}}, "nullable": true}
```
