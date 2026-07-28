---
title: Set DNS preferences
page_id: operation-post-tailnet-tailnet-dns-preferences-79c4f52d
path: operations/dns
description: |-
    Set the DNS preferences for a tailnet; specifically, the MagicDNS setting.
    Note that MagicDNS is dependent on DNS servers.
    Learn about [MagicDNS](https://tailscale.com/kb/1081).

    If there is at least one DNS server, then MagicDNS can be enabled.
    Otherwise, it returns an error.

    Note that removing all nameservers will turn off MagicDNS.
    To reenable it, nameservers must be added back, and MagicDNS must be explicitly turned on.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/dns/preferences
operation_ids:
    - setDnsPreferences
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set DNS preferences

`POST /tailnet/{tailnet}/dns/preferences`

Operation ID: `setDnsPreferences`

Set the DNS preferences for a tailnet; specifically, the MagicDNS setting.
Note that MagicDNS is dependent on DNS servers.
Learn about [MagicDNS](https://tailscale.com/kb/1081).

If there is at least one DNS server, then MagicDNS can be enabled.
Otherwise, it returns an error.

Note that removing all nameservers will turn off MagicDNS.
To reenable it, nameservers must be added back, and MagicDNS must be explicitly turned on.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Set DNS preferences
description: |
    Set the DNS preferences for a tailnet; specifically, the MagicDNS setting.
    Note that MagicDNS is dependent on DNS servers.
    Learn about [MagicDNS](https://tailscale.com/kb/1081).

    If there is at least one DNS server, then MagicDNS can be enabled.
    Otherwise, it returns an error.

    Note that removing all nameservers will turn off MagicDNS.
    To reenable it, nameservers must be added back, and MagicDNS must be explicitly turned on.
operationId: setDnsPreferences
tags:
    - DNS
requestBody:
    content:
        application/json:
            schema:
                $ref: '#/components/schemas/DnsPreferences'
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
