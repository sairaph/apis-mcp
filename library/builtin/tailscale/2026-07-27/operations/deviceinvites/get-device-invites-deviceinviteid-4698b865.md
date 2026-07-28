---
title: Get a device invite
page_id: operation-get-device-invites-deviceinviteid-7f9ac50d
path: operations/deviceinvites
description: |-
    Retrieve a specific device invite.

    OAuth Scope: `device_invites:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /device-invites/{deviceInviteId}
operation_ids:
    - getDeviceInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a device invite

`GET /device-invites/{deviceInviteId}`

Operation ID: `getDeviceInvite`

Retrieve a specific device invite.

OAuth Scope: `device_invites:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceInviteId'
```

## Definition

```yaml
summary: Get a device invite
description: |
    Retrieve a specific device invite.

    OAuth Scope: `device_invites:read`.
operationId: getDeviceInvite
tags:
    - DeviceInvites
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/DeviceInvite'
    '404':
        description: Device invite not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
