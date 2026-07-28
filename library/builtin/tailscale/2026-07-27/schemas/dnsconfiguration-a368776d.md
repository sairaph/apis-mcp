---
title: DnsConfiguration
page_id: schema-dnsconfiguration-a368776d
path: schemas
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DnsConfiguration

```yaml
type: object
properties:
    nameservers:
        type: array
        description: |
            Global DNS resolvers to use. If `preferences.overrideLocalDNS` is true, these override the local OS configuration; otherwise they are used as fallback resolvers.
        items:
            $ref: '#/components/schemas/DnsConfigurationResolver'
        example:
            - address: 8.8.8.8
              useWithExitNode: true
            - address: 1.1.1.1
              useWithExitNode: false
    splitDNS:
        type: object
        description: |
            Map of DNS name suffixes (domains) to lists of resolvers for Split DNS and advanced routing overlays.
        additionalProperties:
            x-additionalPropertiesName: Domain names to DNS resolvers
            type:
                - array
                - 'null'
            items:
                $ref: '#/components/schemas/DnsConfigurationResolver'
        example:
            corp.example.com:
                - address: 10.0.0.53
                  useWithExitNode: true
                - address: 10.0.1.53
                  useWithExitNode: true
            other.internal:
                - address: 10.0.2.53
                  useWithExitNode: false
    searchPaths:
        type: array
        items:
            type: string
        description: |
            Search domain paths to apply.
        example:
            - user1.example.com
            - user2.example.com
    preferences:
        $ref: '#/components/schemas/DnsConfigurationPreferences'
```
