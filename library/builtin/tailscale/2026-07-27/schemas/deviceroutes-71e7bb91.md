---
title: DeviceRoutes
page_id: schema-deviceroutes-71e7bb91
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DeviceRoutes

```yaml
type: object
properties:
    advertisedRoutes:
        type: array
        description: |
            The subnets this device requests to expose.
        items:
            type: string
        example:
            - 10.0.0.0/16
            - 192.168.1.0/24
    enabledRoutes:
        type: array
        description: |
            The subnet routes for this device that have been approved by a tailnet admin.
        items:
            type: string
        example:
            - 10.0.0.0/16
            - 192.168.1.0/24
```
