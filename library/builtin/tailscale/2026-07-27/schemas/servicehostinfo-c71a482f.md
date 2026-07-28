---
title: ServiceHostInfo
page_id: schema-servicehostinfo-c71a482f
path: schemas
description: An information summary for a device hosting a Service.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# ServiceHostInfo

An information summary for a device hosting a Service.

```yaml
type: object
description: |
    An information summary for a device hosting a Service.
properties:
    stableNodeID:
        type: string
        example: n292kg92CNTRL
        description: |
            The preferred identifier for a device.
    approvalLevel:
        type: string
        description: |
            The approval level of the device hosting the Service.
        enum:
            - not-approved
            - approved:auto
            - approved:manual
    configured:
        type: string
        description: The configuration status of the device hosting the Service.
        example: ready
```
