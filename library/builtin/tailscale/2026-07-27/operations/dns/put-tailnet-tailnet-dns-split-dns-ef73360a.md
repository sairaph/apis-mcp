---
title: Set split DNS
page_id: operation-put-tailnet-tailnet-dns-split-dns-2585fe4e
path: operations/dns
description: |-
    Replaces the split DNS settings for a given tailnet.
    Setting the value of a mapping to `null` clears the nameservers for that domain.
    Sending an empty object clears nameservers for all domains.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PUT
api_endpoints:
    - /tailnet/{tailnet}/dns/split-dns
operation_ids:
    - setSplitDns
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set split DNS

`PUT /tailnet/{tailnet}/dns/split-dns`

Operation ID: `setSplitDns`

Replaces the split DNS settings for a given tailnet.
Setting the value of a mapping to `null` clears the nameservers for that domain.
Sending an empty object clears nameservers for all domains.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Set split DNS
description: |
    Replaces the split DNS settings for a given tailnet.
    Setting the value of a mapping to `null` clears the nameservers for that domain.
    Sending an empty object clears nameservers for all domains.
operationId: setSplitDns
tags:
    - DNS
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/SplitDns'
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
