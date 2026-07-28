---
title: Set device name
page_id: operation-post-device-deviceid-name-f659c4c9
path: operations/devices
description: |-
    When a device is added to a tailnet, its Tailscale [device name](https://tailscale.com/kb/1098/machine-names) (also sometimes referred to as machine name) is generated from its OS hostname.
    The device name is the canonical name for the device on your tailnet.

    Device name changes immediately get propogated through your tailnet, so be aware that any existing [Magic DNS](https://tailscale.com/kb/1081/magicdns) URLs using the old name will no longer work.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/name
operation_ids:
    - setDeviceName
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set device name

`POST /device/{deviceId}/name`

Operation ID: `setDeviceName`

When a device is added to a tailnet, its Tailscale [device name](https://tailscale.com/kb/1098/machine-names) (also sometimes referred to as machine name) is generated from its OS hostname.
The device name is the canonical name for the device on your tailnet.

Device name changes immediately get propogated through your tailnet, so be aware that any existing [Magic DNS](https://tailscale.com/kb/1081/magicdns) URLs using the old name will no longer work.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Set device name
description: |
    When a device is added to a tailnet, its Tailscale [device name](https://tailscale.com/kb/1098/machine-names) (also sometimes referred to as machine name) is generated from its OS hostname.
    The device name is the canonical name for the device on your tailnet.

    Device name changes immediately get propogated through your tailnet, so be aware that any existing [Magic DNS](https://tailscale.com/kb/1081/magicdns) URLs using the old name will no longer work.

    OAuth Scope: `devices:core`.
operationId: setDeviceName
tags:
    - Devices
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    name:
                        type: string
                        description: |
                            The new name for the device.

                            This can be provided as either the fully qualified domain name for the device (e.g. "nodename.your-domain.ts.net")
                            or just the base name (e.g. "nodename").

                            If `name` is unset or provided empty, the device's name is reset to be
                            generated from its OS hostname.
                        example: dev-server
                required:
                    - name
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
