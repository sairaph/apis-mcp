---
title: VIPServiceInfoPut
page_id: schema-vipserviceinfoput-8e18b42e
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# VIPServiceInfoPut

```yaml
allOf:
    - type: object
      properties:
        addrs:
            type: array
            description: |
                The IP addresses assigned to the Service.

                - For new Services: either unset or a single IPv4 to assign the Service.
                - For existing Services: an IPv4 and an IPv6. The IPv4 can be updated, but not the IPv6.
            items:
                type: string
            example:
                - 100.93.49.180
                - fd7a:115c:a1e0::3456:3cb4
    - $ref: '#/components/schemas/VIPServiceInfo'
```
