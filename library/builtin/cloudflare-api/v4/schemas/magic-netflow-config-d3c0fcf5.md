---
title: magic_netflow_config
page_id: schema-magic-netflow-config-d3c0fcf5
path: schemas
description: NetFlow configuration for a site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_netflow_config

NetFlow configuration for a site.

```yaml
{"description": "NetFlow configuration for a site.", "type": "object", "properties": {"active_timeout": {"description": "Timeout in seconds for active flows (defaults to 30).", "type": "integer", "example": 30, "maximum": 5400, "minimum": 1}, "collector_ip": {"description": "IPv4 address of the NetFlow collector.", "type": "string", "example": "162.159.65.1"}, "collector_port": {"description": "UDP port of the NetFlow collector (defaults to 2055).", "type": "integer", "example": 2055, "maximum": 65535, "minimum": 1}, "inactive_timeout": {"description": "Timeout in seconds for inactive flows (defaults to 15).", "type": "integer", "example": 15, "maximum": 5400, "minimum": 1}, "sampling_rate": {"description": "Sampling rate for NetFlow records (1 = every packet, 1000 = 1 in 1000 packets). Defaults to 1.", "type": "integer", "example": 100, "maximum": 10000, "minimum": 1}}, "required": ["collector_ip"]}
```
