---
title: Update split DNS
page_id: operation-patch-tailnet-tailnet-dns-split-dns-63ca2659
path: operations/dns
description: |-
    Performs partial updates of the split DNS settings for a given tailnet.
    Only domains specified in the request map will be modified.
    Setting the value of a mapping to `null` clears the nameservers for that domain.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /tailnet/{tailnet}/dns/split-dns
operation_ids:
    - updateSplitDns
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update split DNS

`PATCH /tailnet/{tailnet}/dns/split-dns`

Operation ID: `updateSplitDns`

Performs partial updates of the split DNS settings for a given tailnet.
Only domains specified in the request map will be modified.
Setting the value of a mapping to `null` clears the nameservers for that domain.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Update split DNS
description: |
    Performs partial updates of the split DNS settings for a given tailnet.
    Only domains specified in the request map will be modified.
    Setting the value of a mapping to `null` clears the nameservers for that domain.
operationId: updateSplitDns
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
