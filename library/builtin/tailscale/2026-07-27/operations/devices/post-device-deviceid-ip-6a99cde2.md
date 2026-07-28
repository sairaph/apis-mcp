---
title: Set device IPv4 address
page_id: operation-post-device-deviceid-ip-07fc995b
path: operations/devices
description: |-
    When a device is added to a tailnet, its Tailscale IPv4 address is set at random either from the CGNAT range,
    or a subset of the CGNAT range specified by an [ip pool](https://tailscale.com/kb/1304/ip-pool).
    This endpoint can be used to replace the existing IPv4 address with a specific value.

    This action will break any existing connections to this machine.
    You will need to reconnect to this machine using the new IP address.
    You may also need to flush your DNS cache.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/ip
operation_ids:
    - setDeviceIp
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set device IPv4 address

`POST /device/{deviceId}/ip`

Operation ID: `setDeviceIp`

When a device is added to a tailnet, its Tailscale IPv4 address is set at random either from the CGNAT range,
or a subset of the CGNAT range specified by an [ip pool](https://tailscale.com/kb/1304/ip-pool).
This endpoint can be used to replace the existing IPv4 address with a specific value.

This action will break any existing connections to this machine.
You will need to reconnect to this machine using the new IP address.
You may also need to flush your DNS cache.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Set device IPv4 address
description: |
    When a device is added to a tailnet, its Tailscale IPv4 address is set at random either from the CGNAT range,
    or a subset of the CGNAT range specified by an [ip pool](https://tailscale.com/kb/1304/ip-pool).
    This endpoint can be used to replace the existing IPv4 address with a specific value.

    This action will break any existing connections to this machine.
    You will need to reconnect to this machine using the new IP address.
    You may also need to flush your DNS cache.

    OAuth Scope: `devices:core`.
operationId: setDeviceIp
tags:
    - Devices
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    ipv4:
                        type: string
                        description: |
                            The new IPv4 address for the device.
                        example: 100.80.0.1
                required:
                    - ipv4
responses:
    '200':
        description: Successful operation.
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
