---
title: VIPServiceApproval
page_id: schema-vipserviceapproval-49b0ca48
path: schemas
description: The approval status of a Service on a specific device.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# VIPServiceApproval

The approval status of a Service on a specific device.

```yaml
type: object
description: |
    The approval status of a Service on a specific device.
properties:
    approved:
        type: boolean
        description: |
            Indicates whether the Service is approved on the device.
    autoApproved:
        type: boolean
        description: |
            Indicates whether the Service was auto-approved by an auto-approver.
```
