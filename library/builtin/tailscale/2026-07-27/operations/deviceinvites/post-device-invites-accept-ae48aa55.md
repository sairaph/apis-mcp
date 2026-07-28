---
title: Accept a device invite
page_id: operation-post-device-invites-accept-d2905e87
path: operations/deviceinvites
description: |-
    Accepts the invitation to share a device into the requesting user's tailnet.

    Note that device invites cannot be accepted using an API access token generated from an OAuth client as the shared device is scoped to a user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device-invites/-/accept
operation_ids:
    - acceptDeviceInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Accept a device invite

`POST /device-invites/-/accept`

Operation ID: `acceptDeviceInvite`

Accepts the invitation to share a device into the requesting user's tailnet.

Note that device invites cannot be accepted using an API access token generated from an OAuth client as the shared device is scoped to a user.

## Definition

```yaml
summary: Accept a device invite
description: |
    Accepts the invitation to share a device into the requesting user's tailnet.

    Note that device invites cannot be accepted using an API access token generated from an OAuth client as the shared device is scoped to a user.
operationId: acceptDeviceInvite
tags:
    - DeviceInvites
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    invite:
                        type: string
                        example: https://login.tailscale.com/admin/invite/xxxxxx
                        description: |
                            The URL of the invite (in the form `https://login.tailscale.com/admin/invite/{code}`) or the `{code}` component of the URL.
                required:
                    - invite
            examples:
                usingFullUrl:
                    summary: Using the full invite URL
                    value:
                        invite: https://login.tailscale.com/admin/invite/xxxxxx
                usingCodeComponent:
                    summary: Using `code` component of the invite URL
                    value:
                        invite: xxxxxx
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        device:
                            type: object
                            description: |
                                Information about the device being shared.
                            properties:
                                id:
                                    type: string
                                    example: '12346'
                                    description: |
                                        The `nodeId` for the device.
                                os:
                                    type: string
                                    example: iOS
                                    description: |
                                        The operating system that the device is running.
                                name:
                                    type: string
                                    example: my-phone
                                    description: |
                                        The name of the device.
                                fqdn:
                                    type: string
                                    example: my-phone.something.ts.net
                                    description: |
                                        The MagicDNS name of the device.
                                        Learn more about MagicDNS at https://tailscale.com/kb/1081/.
                                ipv4:
                                    type: string
                                    example: 100.x.y.z
                                    description: |
                                        The IPv4 address of the device.
                                ipv6:
                                    type: string
                                    example: fd7a:115c:x::y:z
                                    description: |
                                        The IPv6 address of the device.
                                includeExitNode:
                                    type: boolean
                                    example: false
                                    description: |
                                        Specifies whether the invited user is able to use the
                                        device as an exit node when the device is advertising as one.
                        sharer:
                            type: object
                            description: |
                                The user who create the device share invite.
                            properties:
                                id:
                                    type: string
                                    example: '22012'
                                    description: |
                                        The ID of the user who created the share invite.
                                displayName:
                                    type: string
                                    example: Some User
                                    description: |
                                        The display name of the user who created the share invite.
                                loginName:
                                    type: string
                                    example: someuser@example.com
                                    description: |
                                        The email address of the user who created the share invite.
                                profilePicURL:
                                    type: string
                                    example: ''
                                    description: |
                                        The profile pic URL for the user who created the share invite.
                        acceptedBy:
                            type: object
                            description: |
                                The user accepting the device share invite.
                            properties:
                                id:
                                    type: string
                                    example: '33233'
                                    description: |
                                        The ID of the user who accepted the share invite.
                                displayName:
                                    type: string
                                    example: Another User
                                    description: |
                                        The display name of the user who accepted the share invite.
                                loginName:
                                    type: string
                                    example: anotheruser@example2.com
                                    description: |
                                        The email address of the user who accepted the share invite.
                                profilePicURL:
                                    type: string
                                    example: ''
                                    description: |
                                        The profile pic URL for the user who accepted the share invite.
    '400':
        $ref: '#/components/responses/400'
    '500':
        $ref: '#/components/responses/500'
```
