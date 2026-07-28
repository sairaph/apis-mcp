---
title: Set device routes
page_id: operation-post-device-deviceid-routes-4ba19650
path: operations/devices
description: |-
    Set a device's enabled subnet routes by replacing the existing list of subnet routes with the supplied parameters.
    [Advertised routes](/kb/1019/subnets#advertise-subnet-routes) cannot be set through the API, since they must be set directly on the device.

    Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
    If a device has advertised routes, they are not exposed to traffic until they are enabled.
    Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

    Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

    OAuth Scope: `devices:routes`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/routes
operation_ids:
    - setDeviceRoutes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set device routes

`POST /device/{deviceId}/routes`

Operation ID: `setDeviceRoutes`

Set a device's enabled subnet routes by replacing the existing list of subnet routes with the supplied parameters.
[Advertised routes](/kb/1019/subnets#advertise-subnet-routes) cannot be set through the API, since they must be set directly on the device.

Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
If a device has advertised routes, they are not exposed to traffic until they are enabled.
Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

OAuth Scope: `devices:routes`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Set device routes
description: |
    Set a device's enabled subnet routes by replacing the existing list of subnet routes with the supplied parameters.
    [Advertised routes](/kb/1019/subnets#advertise-subnet-routes) cannot be set through the API, since they must be set directly on the device.

    Routes must be both advertised and enabled for a device to act as a subnet router or exit node.
    If a device has advertised routes, they are not exposed to traffic until they are enabled.
    Conversely, if routes are enabled before they are advertised, they are not available for routing until the device in question has advertised them.

    Learn more about [subnet routers](/kb/1019/subnets) and [exit nodes](/kb/1103/exit-nodes).

    OAuth Scope: `devices:routes`.
operationId: setDeviceRoutes
tags:
    - Devices
requestBody:
    required: true
    content:
        application/json:
            schema:
                type: object
                properties:
                    routes:
                        type: array
                        description: |
                            The new list of enabled subnet routes.
                        items:
                            type: string
                        example:
                            - 10.0.0.0/16
                            - 192.168.1.0/24
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
