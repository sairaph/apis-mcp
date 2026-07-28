---
title: List device invites
page_id: operation-get-device-deviceid-device-invites-97dd5504
path: operations/deviceinvites
description: |-
    List all share invites for a device.

    OAuth Scope: `device_invites:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /device/{deviceId}/device-invites
operation_ids:
    - listDeviceInvites
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List device invites

`GET /device/{deviceId}/device-invites`

Operation ID: `listDeviceInvites`

List all share invites for a device.

OAuth Scope: `device_invites:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: List device invites
description: |
    List all share invites for a device.

    OAuth Scope: `device_invites:read`.
operationId: listDeviceInvites
tags:
    - DeviceInvites
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: array
                    items:
                        $ref: '#/components/schemas/DeviceInvite'
                    example:
                        - id: '12345'
                          created: '2024-05-08T20:19:51.777861756Z'
                          tailnetId: 59954
                          deviceId: 11055
                          sharerId: 22011
                          allowExitNode: true
                          email: user@example.com
                          lastEmailSentAt: '2024-05-08T20:19:51.777861756Z'
                          inviteUrl: https://login.tailscale.com/admin/invite/<code>
                          accepted: false
                        - id: '12346'
                          created: '2024-04-03T21:38:49.333829261Z'
                          tailnetId: 59954
                          deviceId: 11055
                          sharerId: 22012
                          inviteUrl: https://login.tailscale.com/admin/invite/<code>
                          accepted: true
                          acceptedBy:
                            id: 33223
                            loginName: someone@example.com
                            profilePicUrl: ''
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
