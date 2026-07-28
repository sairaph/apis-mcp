---
title: Delete a user invite
page_id: operation-delete-user-invites-userinviteid-9a1ac220
path: operations/userinvites
description: |-
    Deletes a specific user invite.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - DELETE
api_endpoints:
    - /user-invites/{userInviteId}
operation_ids:
    - deleteUserInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a user invite

`DELETE /user-invites/{userInviteId}`

Operation ID: `deleteUserInvite`

Deletes a specific user invite.

> ⓘ Only permitted for user-owned keys, because invites require an inviting user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userInviteId'
```

## Definition

```yaml
summary: Delete a user invite
description: |
    Deletes a specific user invite.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
operationId: deleteUserInvite
tags:
    - UserInvites
responses:
    '200':
        description: Successful operation.
    '404':
        description: User invite not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
