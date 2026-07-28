---
title: List device routes
page_id: operation-get-device-deviceid-routes-08b2f5f8
path: operations/devices
description: |-
    Retrieve the list of subnet routes that a device is advertising,
    as well as those that are enabled for it.

    Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
    If a device has advertised routes, they are not exposed to traffic until they are enabled.
    Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

    Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

    OAuth Scope: `devices:routes:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /device/{deviceId}/routes
operation_ids:
    - listDeviceRoutes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List device routes

`GET /device/{deviceId}/routes`

Operation ID: `listDeviceRoutes`

Retrieve the list of subnet routes that a device is advertising,
as well as those that are enabled for it.

Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
If a device has advertised routes, they are not exposed to traffic until they are enabled.
Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

OAuth Scope: `devices:routes:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: List device routes
description: |
    Retrieve the list of subnet routes that a device is advertising,
    as well as those that are enabled for it.

    Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
    If a device has advertised routes, they are not exposed to traffic until they are enabled.
    Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

    Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

    OAuth Scope: `devices:routes:read`.
operationId: listDeviceRoutes
tags:
    - Devices
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DeviceRoutes'
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
