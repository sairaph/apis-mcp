---
title: Get DNS preferences
page_id: operation-get-tailnet-tailnet-dns-preferences-b8da08b3
path: operations/dns
description: Retrieves the DNS preferences that are currently set for the given tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/dns/preferences
operation_ids:
    - getDnsPreferences
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get DNS preferences

`GET /tailnet/{tailnet}/dns/preferences`

Operation ID: `getDnsPreferences`

Retrieves the DNS preferences that are currently set for the given tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Get DNS preferences
description: |
    Retrieves the DNS preferences that are currently set for the given tailnet.
operationId: getDnsPreferences
tags:
    - DNS
responses:
    '200':
        description: Succesful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DnsPreferences'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
