---
title: Get a device
page_id: operation-get-device-deviceid-41dc6331
path: operations/devices
description: |-
    Retrieve the details for the specified device.

    OAuth Scope: `devices:core:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /device/{deviceId}
operation_ids:
    - getDevice
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a device

`GET /device/{deviceId}`

Operation ID: `getDevice`

Retrieve the details for the specified device.

OAuth Scope: `devices:core:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Get a device
description: |
    Retrieve the details for the specified device.

    OAuth Scope: `devices:core:read`.
operationId: getDevice
tags:
    - Devices
parameters:
    - $ref: '#/components/parameters/fields'
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/Device'
    '400':
        description: Invalid ID supplied.
        $ref: '#/components/responses/400'
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
