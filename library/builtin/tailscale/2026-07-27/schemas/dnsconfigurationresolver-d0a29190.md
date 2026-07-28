---
title: DnsConfigurationResolver
page_id: schema-dnsconfigurationresolver-d0a29190
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DnsConfigurationResolver

```yaml
type: object
properties:
    address:
        type: string
        description: |
            IPv4 or IPv6 address of the DNS resolver.
        example: 1.1.1.1
    useWithExitNode:
        type: boolean
        description: |
            If true, this resolver should still be used when a device is configured to use a Tailscale exit node. Requires Tailscale v1.88.1 or later.
        example: true
```
