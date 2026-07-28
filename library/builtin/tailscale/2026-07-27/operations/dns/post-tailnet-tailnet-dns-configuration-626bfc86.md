---
title: Set DNS configuration
page_id: operation-post-tailnet-tailnet-dns-configuration-26c98c65
path: operations/dns
description: |-
    Replaces the DNS configuration for the given tailnet.

    - `nameservers` defines the global resolvers to use when `preferences.overrideLocalDNS` is true.
    - `splitDNS` maps DNS name suffixes (domains) to lists of resolvers for Split DNS.
    - `searchPaths` sets custom DNS search domain paths.
    - `preferences.overrideLocalDNS` controls whether resolvers in `nameservers` override the local OS configuration (true) or are used or local resolvers are used (false). Defaults to false.
    - `preferences.magicDNS` enables MagicDNS. Defaults to false.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/dns/configuration
operation_ids:
    - setDnsConfiguration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set DNS configuration

`POST /tailnet/{tailnet}/dns/configuration`

Operation ID: `setDnsConfiguration`

Replaces the DNS configuration for the given tailnet.

- `nameservers` defines the global resolvers to use when `preferences.overrideLocalDNS` is true.
- `splitDNS` maps DNS name suffixes (domains) to lists of resolvers for Split DNS.
- `searchPaths` sets custom DNS search domain paths.
- `preferences.overrideLocalDNS` controls whether resolvers in `nameservers` override the local OS configuration (true) or are used or local resolvers are used (false). Defaults to false.
- `preferences.magicDNS` enables MagicDNS. Defaults to false.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Set DNS configuration
description: |
    Replaces the DNS configuration for the given tailnet.

    - `nameservers` defines the global resolvers to use when `preferences.overrideLocalDNS` is true.
    - `splitDNS` maps DNS name suffixes (domains) to lists of resolvers for Split DNS.
    - `searchPaths` sets custom DNS search domain paths.
    - `preferences.overrideLocalDNS` controls whether resolvers in `nameservers` override the local OS configuration (true) or are used or local resolvers are used (false). Defaults to false.
    - `preferences.magicDNS` enables MagicDNS. Defaults to false.
operationId: setDnsConfiguration
tags:
    - DNS
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/DnsConfiguration'
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DnsConfiguration'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
