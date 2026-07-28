---
title: Set DNS search paths
page_id: operation-post-tailnet-tailnet-dns-searchpaths-574e71d6
path: operations/dns
description: Replaces the list of search paths for the given tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/dns/searchpaths
operation_ids:
    - setDnsSearchPaths
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set DNS search paths

`POST /tailnet/{tailnet}/dns/searchpaths`

Operation ID: `setDnsSearchPaths`

Replaces the list of search paths for the given tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Set DNS search paths
description: |
    Replaces the list of search paths for the given tailnet.
operationId: setDnsSearchPaths
tags:
    - DNS
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/DnsSearchPaths'
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
