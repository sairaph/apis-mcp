---
title: DnsConfigurationPreferences
page_id: schema-dnsconfigurationpreferences-475cab58
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DnsConfigurationPreferences

```yaml
type: object
properties:
    overrideLocalDNS:
        type: boolean
        description: |
            If true, resolvers in `nameservers` override the local OS DNS configuration; if false, local resolvers are used.
        example: true
    magicDNS:
        type: boolean
        description: |
            Whether MagicDNS is enabled for this tailnet.
        example: true
```
