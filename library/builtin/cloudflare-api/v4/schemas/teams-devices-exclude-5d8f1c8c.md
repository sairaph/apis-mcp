---
title: teams-devices_exclude
page_id: schema-teams-devices-exclude-5d8f1c8c
path: schemas
description: List of routes excluded in the WARP client's tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_exclude

List of routes excluded in the WARP client's tunnel.

```yaml
{"description": "List of routes excluded in the WARP client's tunnel.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_split_tunnel"}, "default": [{"address": "10.0.0.0/8"}, {"address": "100.64.0.0/10"}, {"address": "169.254.0.0/16", "description": "DHCP Unspecified"}, {"address": "172.16.0.0/12"}, {"address": "192.0.0.0/24"}, {"address": "192.168.0.0/16"}, {"address": "224.0.0.0/24"}, {"address": "240.0.0.0/4"}, {"address": "255.255.255.255/32", "description": "DHCP Broadcast"}, {"address": "fe80::/10", "description": "IPv6 Link Local"}, {"address": "fd00::/8"}, {"address": "ff01::/16"}, {"address": "ff02::/16"}, {"address": "ff03::/16"}, {"address": "ff04::/16"}, {"address": "ff05::/16"}]}
```
