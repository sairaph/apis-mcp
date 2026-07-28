---
title: Get split DNS
page_id: operation-get-tailnet-tailnet-dns-split-dns-97aca5cc
path: operations/dns
description: Retrieves the split DNS settings, which is a map from domains to lists of nameservers, that is currently set for the given tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/dns/split-dns
operation_ids:
    - getSplitDns
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get split DNS

`GET /tailnet/{tailnet}/dns/split-dns`

Operation ID: `getSplitDns`

Retrieves the split DNS settings, which is a map from domains to lists of nameservers, that is currently set for the given tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Get split DNS
description: |
    Retrieves the split DNS settings, which is a map from domains to lists of nameservers, that is currently set for the given tailnet.
operationId: getSplitDns
tags:
    - DNS
responses:
    '200':
        description: Succesful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/SplitDns'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
