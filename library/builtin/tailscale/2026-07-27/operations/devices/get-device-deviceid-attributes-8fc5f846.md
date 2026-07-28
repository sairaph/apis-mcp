---
title: Get device posture attributes
page_id: operation-get-device-deviceid-attributes-1e20a1a6
path: operations/devices
description: |-
    Retrieve all posture attributes for the specified device.
    This returns a JSON object of all the key-value pairs of posture attributes for the device.

    OAuth Scope: `devices:posture_attributes:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /device/{deviceId}/attributes
operation_ids:
    - getDevicePostureAttributes
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get device posture attributes

`GET /device/{deviceId}/attributes`

Operation ID: `getDevicePostureAttributes`

Retrieve all posture attributes for the specified device.
This returns a JSON object of all the key-value pairs of posture attributes for the device.

OAuth Scope: `devices:posture_attributes:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Get device posture attributes
description: |
    Retrieve all posture attributes for the specified device.
    This returns a JSON object of all the key-value pairs of posture attributes for the device.

    OAuth Scope: `devices:posture_attributes:read`.
operationId: getDevicePostureAttributes
tags:
    - Devices
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DevicePostureAttributes'
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
