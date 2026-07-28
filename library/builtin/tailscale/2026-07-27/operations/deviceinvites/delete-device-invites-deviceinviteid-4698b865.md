---
title: Delete a device invite
page_id: operation-delete-device-invites-deviceinviteid-7c5f1619
path: operations/deviceinvites
description: |-
    Delete a specific device invite.

    OAuth Scope: `device_invites`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /device-invites/{deviceInviteId}
operation_ids:
    - deleteDeviceInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a device invite

`DELETE /device-invites/{deviceInviteId}`

Operation ID: `deleteDeviceInvite`

Delete a specific device invite.

OAuth Scope: `device_invites`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceInviteId'
```

## Definition

```yaml
summary: Delete a device invite
description: |
    Delete a specific device invite.

    OAuth Scope: `device_invites`.
operationId: deleteDeviceInvite
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
