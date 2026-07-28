---
title: Authorize device
page_id: operation-post-device-deviceid-authorized-970b45f2
path: operations/devices
description: |-
    This call marks a device as authorized or revokes its authorization for tailnets where device authorization is required,
    according to the authorized field in the payload.

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/authorized
operation_ids:
    - authorizeDevice
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Authorize device

`POST /device/{deviceId}/authorized`

Operation ID: `authorizeDevice`

This call marks a device as authorized or revokes its authorization for tailnets where device authorization is required,
according to the authorized field in the payload.

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Authorize device
description: |
    This call marks a device as authorized or revokes its authorization for tailnets where device authorization is required,
    according to the authorized field in the payload.

    OAuth Scope: `devices:core`.
operationId: authorizeDevice
tags:
    - Devices
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    authorized:
                        type: boolean
                        description: |
                            - If `true`, authorize a new device or re-authorize a previously deauthorized device.
                            - If `false`, deauthorize an authorized device.
                required:
                    - authorized
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
