---
title: SplitDns
page_id: schema-splitdns-72ddd838
path: schemas
description: Map of domain names to lists of nameservers or to `null`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# SplitDns

Map of domain names to lists of nameservers or to `null`.

```yaml
type: object
additionalProperties:
    x-additionalPropertiesName: Domain names to DNS
    type:
        - array
        - 'null'
    items:
        type: string
example:
    example.com:
        - 1.1.1.1
        - 1.2.3.4
    other.com:
        - 2.2.2.2
description: |
    Map of domain names to lists of nameservers or to `null`.
```
