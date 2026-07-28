---
title: UserInvite
page_id: schema-userinvite-38303e15
path: schemas
description: |-
    A user invite is an active invitation that lets a user join a tailnet
    with a preassigned [user role](https://tailscale.com/kb/1138/user-roles).

    Each user invite has a unique ID that is used to identify the invite
    in API calls. You can find all user invite IDs for a particular tailnet
    by [listing user invites](#tag/userinvites/get/tailnet/{tailnet}/user-invites).
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# UserInvite

A user invite is an active invitation that lets a user join a tailnet
with a preassigned [user role](https://tailscale.com/kb/1138/user-roles).

Each user invite has a unique ID that is used to identify the invite
in API calls. You can find all user invite IDs for a particular tailnet
by [listing user invites](#tag/userinvites/get/tailnet/{tailnet}/user-invites).

```yaml
type: object
description: |
    A user invite is an active invitation that lets a user join a tailnet
    with a preassigned [user role](https://tailscale.com/kb/1138/user-roles).

    Each user invite has a unique ID that is used to identify the invite
    in API calls. You can find all user invite IDs for a particular tailnet
    by [listing user invites](#tag/userinvites/get/tailnet/{tailnet}/user-invites).
properties:
    id:
        type: string
        example: '12346'
        description: |
            The unique identifier for the invite.
            Supply this value wherever `userInviteId` is indicated in the endpoint.
    role:
        type: string
        enum:
            - member
            - admin
            - it-admin
            - network-admin
            - billing-admin
            - auditor
        example: admin
        description: |
            The tailnet user role to assign to the invited user upon accepting the invite.
    tailnetId:
        type: integer
        format: int64
        example: 59954
        description: |
            The ID of the tailnet to which the user was invited.
    inviterId:
        type: integer
        format: int64
        example: 22012
        description: |
            The ID of the user who created the invite.
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
            Included when `email` is not part of the tailnet's domain,
            or when `email` is empty. It is the link to accept the invite.

            When included, anyone with this link can accept the invite.
            It is not restricted to the person to which the invite was emailed.

            When `email` is part of the tailnet's domain (has the same @domain.com
            suffix as the tailnet), the user can join the tailnet automatically by
            logging in with their domain email at https://login.tailscale.com/start.
            They'll be assigned the specified `role` upon signing in for the first
            time.
required:
    - id
    - role
    - tailnetId
    - inviterId
```
