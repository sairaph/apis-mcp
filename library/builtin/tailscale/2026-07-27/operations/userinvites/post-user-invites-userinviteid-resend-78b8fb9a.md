---
title: Resend a user invite
page_id: operation-post-user-invites-userinviteid-resend-cb7499e5
path: operations/userinvites
description: |-
    Resend a user invite by email. You can only use this if the specified invite
    was originally created with an email specified.
    Refer to [creating user invites for a tailnet](#tag/userinvites/post/tailnet/{tailnet}/user-invites).

    Note: Invite resends are rate limited to one per minute.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /user-invites/{userInviteId}/resend
operation_ids:
    - resendUserInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Resend a user invite

`POST /user-invites/{userInviteId}/resend`

Operation ID: `resendUserInvite`

Resend a user invite by email. You can only use this if the specified invite
was originally created with an email specified.
Refer to [creating user invites for a tailnet](#tag/userinvites/post/tailnet/{tailnet}/user-invites).

Note: Invite resends are rate limited to one per minute.

> ⓘ Only permitted for user-owned keys, because invites require an inviting user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userInviteId'
```

## Definition

```yaml
summary: Resend a user invite
description: |
    Resend a user invite by email. You can only use this if the specified invite
    was originally created with an email specified.
    Refer to [creating user invites for a tailnet](#tag/userinvites/post/tailnet/{tailnet}/user-invites).

    Note: Invite resends are rate limited to one per minute.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
operationId: resendUserInvite
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
