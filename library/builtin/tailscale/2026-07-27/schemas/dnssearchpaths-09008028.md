---
title: DnsSearchPaths
page_id: schema-dnssearchpaths-09008028
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DnsSearchPaths

```yaml
type: object
properties:
    searchPaths:
        type: array
        items:
            type: string
        example:
            - user1.example.com
            - user2.example.com
        description: |
            The search domains for the given tailnet.
required:
    - searchPaths
```
