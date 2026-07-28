---
title: Expire a device's key
page_id: operation-post-device-deviceid-expire-b987a832
path: operations/devices
description: |-
    Mark a device's node key as expired.
    This will require the device to re-authenticate in order to connect to the tailnet.
    The device must belong to the requesting user's tailnet.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/expire
operation_ids:
    - expireDeviceKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Expire a device's key

`POST /device/{deviceId}/expire`

Operation ID: `expireDeviceKey`

Mark a device's node key as expired.
This will require the device to re-authenticate in order to connect to the tailnet.
The device must belong to the requesting user's tailnet.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Expire a device's key
description: |
    Mark a device's node key as expired.
    This will require the device to re-authenticate in order to connect to the tailnet.
    The device must belong to the requesting user's tailnet.

    OAuth Scope: `devices:core`.
operationId: expireDeviceKey
tags:
    - Devices
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
