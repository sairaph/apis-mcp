---
title: Set DNS nameservers
page_id: operation-post-tailnet-tailnet-dns-nameservers-0706ae23
path: operations/dns
description: |-
    Replaces the list of global DNS nameservers for the given tailnet with the list supplied in the request.

    Note that changing the list of DNS nameservers may also affect the status of MagicDNS (if MagicDNS is on; learn about [MagicDNS](https://tailscale.com/kb/1081)).
    If all nameservers have been removed, MagicDNS will be automatically disabled (until explicitly turned back on by the user).
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/dns/nameservers
operation_ids:
    - setDnsNameservers
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set DNS nameservers

`POST /tailnet/{tailnet}/dns/nameservers`

Operation ID: `setDnsNameservers`

Replaces the list of global DNS nameservers for the given tailnet with the list supplied in the request.

Note that changing the list of DNS nameservers may also affect the status of MagicDNS (if MagicDNS is on; learn about [MagicDNS](https://tailscale.com/kb/1081)).
If all nameservers have been removed, MagicDNS will be automatically disabled (until explicitly turned back on by the user).

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Set DNS nameservers
description: |
    Replaces the list of global DNS nameservers for the given tailnet with the list supplied in the request.

    Note that changing the list of DNS nameservers may also affect the status of MagicDNS (if MagicDNS is on; learn about [MagicDNS](https://tailscale.com/kb/1081)).
    If all nameservers have been removed, MagicDNS will be automatically disabled (until explicitly turned back on by the user).
operationId: setDnsNameservers
tags:
    - DNS
requestBody:
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
responses:
    '200':
        description: Succesful operation.
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
                        magicDNS:
                            type: boolean
                            example: true
                            description: |
                                Whether MagicDNS is active for this tailnet.
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
