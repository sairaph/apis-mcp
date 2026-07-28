---
title: Get DNS configuration
page_id: operation-get-tailnet-tailnet-dns-configuration-4223e14c
path: operations/dns
description: Retrieves the full DNS configuration for a tailnet, including global nameservers, split DNS routes, search paths, and MagicDNS configuration.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/dns/configuration
operation_ids:
    - getDnsConfiguration
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get DNS configuration

`GET /tailnet/{tailnet}/dns/configuration`

Operation ID: `getDnsConfiguration`

Retrieves the full DNS configuration for a tailnet, including global nameservers, split DNS routes, search paths, and MagicDNS configuration.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Get DNS configuration
description: |
    Retrieves the full DNS configuration for a tailnet, including global nameservers, split DNS routes, search paths, and MagicDNS configuration.
operationId: getDnsConfiguration
tags:
    - DNS
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
