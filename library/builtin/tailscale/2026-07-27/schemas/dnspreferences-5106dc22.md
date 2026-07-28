---
title: DnsPreferences
page_id: schema-dnspreferences-5106dc22
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DnsPreferences

```yaml
type: object
properties:
    magicDNS:
        type: boolean
        example: true
        description: |
            Whether MagicDNS is active for this tailnet.
required:
    - magicDNS
```
