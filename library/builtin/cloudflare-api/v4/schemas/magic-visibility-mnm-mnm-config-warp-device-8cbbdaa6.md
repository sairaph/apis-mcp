---
title: magic-visibility-mnm_mnm_config_warp_device
page_id: schema-magic-visibility-mnm-mnm-config-warp-device-8cbbdaa6
path: schemas
description: Object representing a warp device with an ID and name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-mnm_mnm_config_warp_device

Object representing a warp device with an ID and name.

```yaml
{"description": "Object representing a warp device with an ID and name.", "type": "object", "properties": {"id": {"description": "Unique identifier for the warp device.", "type": "string", "example": "5360368d-b351-4791-abe1-93550dabd351", "x-auditable": true}, "name": {"description": "Name of the warp device.", "type": "string", "example": "My warp device", "x-auditable": true}, "router_ip": {"description": "IPv4 CIDR of the router sourcing flow data associated with this warp device. Only /32 addresses are currently supported.", "type": "string", "example": "203.0.113.1", "x-auditable": true}}, "required": ["id", "name", "router_ip"]}
```
