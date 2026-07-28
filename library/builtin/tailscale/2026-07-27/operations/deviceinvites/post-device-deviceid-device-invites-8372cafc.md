---
title: Create device invites
page_id: operation-post-device-deviceid-device-invites-0ff745a7
path: operations/deviceinvites
description: |-
    Create new share invites for a device.

    Note that device invites cannot be created using an API access token generated from an OAuth client as the shared device is scoped to a user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/device-invites
operation_ids:
    - createDeviceInvites
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create device invites

`POST /device/{deviceId}/device-invites`

Operation ID: `createDeviceInvites`

Create new share invites for a device.

Note that device invites cannot be created using an API access token generated from an OAuth client as the shared device is scoped to a user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Create device invites
description: |
    Create new share invites for a device.

    Note that device invites cannot be created using an API access token generated from an OAuth client as the shared device is scoped to a user.
operationId: createDeviceInvites
tags:
    - DeviceInvites
requestBody:
    description: Device invites to create.
    content:
        application/json:
            schema:
                type: array
                items:
                    type: object
                    properties:
                        multiUse:
                            type: boolean
                            example: false
                            description: |
                                Whether the invite can be accepted more than once.
                                When set to `true`, it results in an invite that can be accepted up to 1,000 times.
                        allowExitNode:
                            type: boolean
                            example: false
                            description: |
                                Whether the invited user can use the device as an exit node when it advertises as one.
                        email:
                            type: string
                            example: user@example.com
                            description: |
                                The email to send the created invite to.
                                If not set, the endpoint generates and returns an invite URL (but doesn't send it out).
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
                          accepted: false
    '404':
        description: Device not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
