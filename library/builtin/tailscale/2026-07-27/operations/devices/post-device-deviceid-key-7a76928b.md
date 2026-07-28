---
title: Update device key
page_id: operation-post-device-deviceid-key-736902d1
path: operations/devices
description: |-
    When a device is added to a tailnet, its key expiry is set according to the tailnet's key expiry setting.
    If the key is not refreshed and expires, the device can no longer communicate with other devices in the tailnet.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/key
operation_ids:
    - updateDeviceKey
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update device key

`POST /device/{deviceId}/key`

Operation ID: `updateDeviceKey`

When a device is added to a tailnet, its key expiry is set according to the tailnet's key expiry setting.
If the key is not refreshed and expires, the device can no longer communicate with other devices in the tailnet.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Update device key
description: |
    When a device is added to a tailnet, its key expiry is set according to the tailnet's key expiry setting.
    If the key is not refreshed and expires, the device can no longer communicate with other devices in the tailnet.

    OAuth Scope: `devices:core`.
operationId: updateDeviceKey
tags:
    - Devices
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    keyExpiryDisabled:
                        type: boolean
                        example: true
                        description: |
                            - If `true`, disable the device's key expiry. The original key expiry time is still maintained. Upon re-enabling, the key will expire at that original time.
                            - If `false`, enable the device's key expiry. Sets the key to expire at the original expiry time prior to disabling. The key may already have expired. In that case, the device must be re-authenticated.
                required:
                    - keyExpiryDisabled
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
