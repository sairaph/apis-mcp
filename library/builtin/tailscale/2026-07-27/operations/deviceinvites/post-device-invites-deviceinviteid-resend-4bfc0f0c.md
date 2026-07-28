---
title: Resend a device invite
page_id: operation-post-device-invites-deviceinviteid-resend-b176c752
path: operations/deviceinvites
description: |-
    Resend a device invite by email. You can only use this if the specified invite
    was originally created with an email specified.
    Refer to [creating device invites for a device](#tag/deviceinvites/post/device/{deviceId}/device-invites).

    Note: Invite resends are rate limited to one per minute.

    Note that device invites cannot be resent using an API access token generated from an OAuth client as the shared device is scoped to a user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device-invites/{deviceInviteId}/resend
operation_ids:
    - resendDeviceInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Resend a device invite

`POST /device-invites/{deviceInviteId}/resend`

Operation ID: `resendDeviceInvite`

Resend a device invite by email. You can only use this if the specified invite
was originally created with an email specified.
Refer to [creating device invites for a device](#tag/deviceinvites/post/device/{deviceId}/device-invites).

Note: Invite resends are rate limited to one per minute.

Note that device invites cannot be resent using an API access token generated from an OAuth client as the shared device is scoped to a user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceInviteId'
```

## Definition

```yaml
summary: Resend a device invite
description: |
    Resend a device invite by email. You can only use this if the specified invite
    was originally created with an email specified.
    Refer to [creating device invites for a device](#tag/deviceinvites/post/device/{deviceId}/device-invites).

    Note: Invite resends are rate limited to one per minute.

    Note that device invites cannot be resent using an API access token generated from an OAuth client as the shared device is scoped to a user.
operationId: resendDeviceInvite
tags:
    - DeviceInvites
responses:
    '200':
        description: Successful operation.
    '404':
        description: Device invite not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
