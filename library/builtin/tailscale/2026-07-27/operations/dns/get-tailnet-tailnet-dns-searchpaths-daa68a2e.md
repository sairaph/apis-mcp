---
title: List DNS search paths
page_id: operation-get-tailnet-tailnet-dns-searchpaths-d06df506
path: operations/dns
description: Retrieves the list of search paths, also referred to as *search domains*, that is currently set for the given tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/dns/searchpaths
operation_ids:
    - listDnsSearchPaths
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List DNS search paths

`GET /tailnet/{tailnet}/dns/searchpaths`

Operation ID: `listDnsSearchPaths`

Retrieves the list of search paths, also referred to as *search domains*, that is currently set for the given tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List DNS search paths
description: |
    Retrieves the list of search paths, also referred to as *search domains*, that is currently set for the given tailnet.
operationId: listDnsSearchPaths
tags:
    - DNS
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DnsSearchPaths'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
