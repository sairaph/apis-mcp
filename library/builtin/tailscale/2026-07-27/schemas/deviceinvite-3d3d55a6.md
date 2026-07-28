---
title: DeviceInvite
page_id: schema-deviceinvite-3d3d55a6
path: schemas
description: |-
    A device invite is an invitation that shares a device with an external
    user (a user not in the device's tailnet).

    Each device invite has a unique ID that is used to identify the invite
    in API calls. You can find all device invite IDs for a particular device
    by [listing all device invites](#tag/deviceinvites/POST/device/{deviceId}/device-invites)
    for a device.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# DeviceInvite

A device invite is an invitation that shares a device with an external
user (a user not in the device's tailnet).

Each device invite has a unique ID that is used to identify the invite
in API calls. You can find all device invite IDs for a particular device
by [listing all device invites](#tag/deviceinvites/POST/device/{deviceId}/device-invites)
for a device.

```yaml
type: object
description: |
    A device invite is an invitation that shares a device with an external
    user (a user not in the device's tailnet).

    Each device invite has a unique ID that is used to identify the invite
    in API calls. You can find all device invite IDs for a particular device
    by [listing all device invites](#tag/deviceinvites/POST/device/{deviceId}/device-invites)
    for a device.
properties:
    id:
        type: string
        example: '12346'
        description: |
            The unique identifier for the invite.
            Supply this value wherever {deviceInviteId} is indicated in the endpoint.
    created:
        type: string
        format: date-time
        example: '2024-04-03T21:38:49.333829261Z'
        description: |
            The creation time of the invite.
    tailnetId:
        type: integer
        format: int64
        example: 59954
        description: |
            The ID of the tailnet to which the shared device belongs.
    deviceId:
        type: integer
        format: int64
        example: 11055
        description: |
            The ID of the device being shared.
    sharerId:
        type: integer
        format: int64
        example: 22012
        description: |
            The ID of the user who created the share invite.
    multiUse:
        type: boolean
        example: false
        description: |
            Specifies whether this device invite can be accepted
            more than once.
    allowExitNode:
        type: boolean
        example: false
        description: |
            Specifies whether the invited user is able to use the
            device as an exit node when the device is advertising as one.
    email:
        type: string
        example: user@example.com
        description: |
            The email to which the invite was sent.
            If empty, the invite was not emailed to anyone,
            but the inviteUrl can be shared manually.
    lastEmailSentAt:
        type: string
        format: date-time
        example: '2024-04-03T21:38:49.333829261Z'
        description: |
            The last time the invite was attempted to be sent to Email.
            Only ever set if `email` is not empty.
    inviteUrl:
        type: string
        example: https://login.tailscale.com/admin/invite/<code>
        description: |
            The link to accept the invite.
            Anyone with this link can accept the invite.
            It is not restricted to the person to which the invite was emailed.
    accepted:
        type: boolean
        example: false
        description: |
            `true` when the share invite has been accepted.
    acceptedBy:
        type: object
        description: |
            Set when the invite has been accepted.
            It holds information about the user who accepted the share invite.
        properties:
            id:
                type: integer
                format: int64
                example: 33223
                description: |
                    The ID of the user who accepted the share invite.
            loginName:
                type: string
                example: someone@example.com
                description: |
                    The login name of the user who accepted the share invite.
            profilePicUrl:
                type: string
                example: ''
                description: |
                    The profile pic URL for the user who accepted the share invite.
```
