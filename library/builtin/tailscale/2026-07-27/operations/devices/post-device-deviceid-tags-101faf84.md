---
title: Set device tags
page_id: operation-post-device-deviceid-tags-12c24790
path: operations/devices
description: |-
    Tags let you assign an identity to a device that is separate from human users, and use that identity as part of an ACL to restrict access.
    Tags are similar to role accounts, but more flexible.

    Tags are created in the tailnet policy file by defining the tag and an owner of the tag.
    Once a device is tagged, the tag is the owner of that device.
    A single node can have multiple tags assigned.

    Consult the policy file for your tailnet in the [admin console](https://login.tailscale.com/admin/acls) for the list of tags that have been created for your tailnet.
    Learn more about [tags](https://tailscale.com/kb/1068/).

    OAuth Scope: `devices:core`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /device/{deviceId}/tags
operation_ids:
    - setDeviceTags
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Set device tags

`POST /device/{deviceId}/tags`

Operation ID: `setDeviceTags`

Tags let you assign an identity to a device that is separate from human users, and use that identity as part of an ACL to restrict access.
Tags are similar to role accounts, but more flexible.

Tags are created in the tailnet policy file by defining the tag and an owner of the tag.
Once a device is tagged, the tag is the owner of that device.
A single node can have multiple tags assigned.

Consult the policy file for your tailnet in the [admin console](https://login.tailscale.com/admin/acls) for the list of tags that have been created for your tailnet.
Learn more about [tags](https://tailscale.com/kb/1068/).

OAuth Scope: `devices:core`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/deviceId'
```

## Definition

```yaml
summary: Set device tags
description: |
    Tags let you assign an identity to a device that is separate from human users, and use that identity as part of an ACL to restrict access.
    Tags are similar to role accounts, but more flexible.

    Tags are created in the tailnet policy file by defining the tag and an owner of the tag.
    Once a device is tagged, the tag is the owner of that device.
    A single node can have multiple tags assigned.

    Consult the policy file for your tailnet in the [admin console](https://login.tailscale.com/admin/acls) for the list of tags that have been created for your tailnet.
    Learn more about [tags](https://tailscale.com/kb/1068/).

    OAuth Scope: `devices:core`.
operationId: setDeviceTags
tags:
    - Devices
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    tags:
                        type: array
                        description: |
                            The new list of tags for the device.
                        items:
                            type: string
                        example:
                            - tag:foo
                            - tag:bar
responses:
    '200':
        description: Successful operation.
    '400':
        $ref: '#/components/responses/400'
    '500':
        $ref: '#/components/responses/500'
    '504':
        $ref: '#/components/responses/504'
```
