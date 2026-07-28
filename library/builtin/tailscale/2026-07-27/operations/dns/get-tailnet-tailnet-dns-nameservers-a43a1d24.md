---
title: List DNS nameservers
page_id: operation-get-tailnet-tailnet-dns-nameservers-df59731a
path: operations/dns
description: Lists the global DNS nameservers for a tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/dns/nameservers
operation_ids:
    - listDnsNameservers
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List DNS nameservers

`GET /tailnet/{tailnet}/dns/nameservers`

Operation ID: `listDnsNameservers`

Lists the global DNS nameservers for a tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List DNS nameservers
description: |
    Lists the global DNS nameservers for a tailnet.
operationId: listDnsNameservers
tags:
    - DNS
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        dns:
                            type: array
                            items:
                                type: string
                            example:
                                - 8.8.8.8
                                - 1.2.3.4
                            description: |
                                DNS nameservers.
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
