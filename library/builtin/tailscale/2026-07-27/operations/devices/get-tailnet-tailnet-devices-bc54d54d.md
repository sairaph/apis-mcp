---
title: List tailnet devices
page_id: operation-get-tailnet-tailnet-devices-1cb1c7ba
path: operations/devices
description: |-
    Lists the devices in a tailnet.

    OAuth Scope: `devices:core:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/devices
operation_ids:
    - listTailnetDevices
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List tailnet devices

`GET /tailnet/{tailnet}/devices`

Operation ID: `listTailnetDevices`

Lists the devices in a tailnet.

OAuth Scope: `devices:core:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/fields'
- in: query
  name: <field>=<value> filters
  description: |
    This endpoint supports server-side filtering of devices by specifying one
    or more filters in the form `<field>=<value>`. Fields must be a top-level
    device property - e.g. `isEphemeral`, `tags`, `hostname`, etc. Values are
    matched as _exact_ matches. Properties with simple types (strings, numbers,
    dates, etc) and lists are supported. Properties that are complex objects
    (e.g. `clientConnectivity`) are not supported. When multiple parameters are
    provided, the server performs a logical `AND` across all filter parameters
    before returning results. For example,
    `isEphemeral=true&tags=tag:prod&tags=tag:subnetrouter` would return devices
    where `isEphemeral` is `true` and `tags` contains both `tag:prod` and
    `tag:subnetrouter`.
  schema:
    type: string
```

## Definition

```yaml
summary: List tailnet devices
description: |
    Lists the devices in a tailnet.

    OAuth Scope: `devices:core:read`.
operationId: listTailnetDevices
tags:
    - Devices
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        devices:
                            type: array
                            items:
                                $ref: '#/components/schemas/Device'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        description: Request took too long to process, please try again later.
        $ref: '#/components/responses/504'
```
