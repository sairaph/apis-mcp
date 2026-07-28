---
title: Delete a device
page_id: operation-delete-device-deviceid-906aa08e
path: operations/devices
description: |-
    Deletes the device from its tailnet.
    The device must belong to the requesting user's tailnet.
    Deleting devices shared with the tailnet is not supported.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /device/{deviceId}
operation_ids:
    - deleteDevice
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a device

`DELETE /device/{deviceId}`

Operation ID: `deleteDevice`

Deletes the device from its tailnet.
The device must belong to the requesting user's tailnet.
Deleting devices shared with the tailnet is not supported.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Delete a device
description: |
    Deletes the device from its tailnet.
    The device must belong to the requesting user's tailnet.
    Deleting devices shared with the tailnet is not supported.

    OAuth Scope: `devices:core`.
operationId: deleteDevice
tags:
    - Devices
responses:
    '200':
        description: Successful operation.
    '400':
        description: Invalid device value.
        $ref: '#/components/responses/400'
    '500':
        $ref: '#/components/responses/500'
    '501':
        description: Device not owned by tailnet.
        $ref: '#/components/responses/501'
    '504':
        $ref: '#/components/responses/504'
```
