---
title: User
page_id: schema-user-b512d97e
path: schemas
description: Representation of a user within a tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# User

Representation of a user within a tailnet.

```yaml
type: object
description: |
    Representation of a user within a tailnet.
properties:
    id:
        type: string
        example: '123456'
        description: |
            The unique identifier for the user.
            Supply this value wherever {userId} is indicated in an endpoint.
    displayName:
        type: string
        example: Some User
        description: |
            The name of the user.
    loginName:
        type: string
        example: someuser@example.com
        description: |
            The emailish login name of the user.
    profilePicUrl:
        type: string
        example: ''
        description: |
            The profile pic URL for the user.
    tailnetId:
        type: string
        example: example.com
        description: |
            The tailnet that owns the user.
    created:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The time the user joined their tailnet.
    type:
        type: string
        enum:
            - member
            - shared
        example: member
        description: |
            The type of relation this user has to the tailnet associated with the request.
    role:
        type: string
        enum:
            - owner
            - member
            - admin
            - it-admin
            - network-admin
            - billing-admin
            - auditor
        example: member
        description: |
            The role of the user. Learn more about [user roles](kb/1138/user-roles).
    status:
        type: string
        enum:
            - active
            - idle
            - suspended
            - needs-approval
            - over-billing-limit
        x-enumDescriptions:
            active: Last seen within 28 days.
            idle: Last seen longer than 28 days.
            suspended: Suspended from accessing the tailnet.
            needs-approval: Unable to join tailnet until approved.
            over-billing-limit: Unable to join tailnet until billing count increased.
        example: active
        description: |
            The status of the user.
    deviceCount:
        type: integer
        example: 4
        description: |
            Number of devices the user owns.
    lastSeen:
        type: string
        format: date-time
        example: '2022-12-01T05:23:30Z'
        description: |
            The later of either:
            - The last time any of the user's nodes were connected to the network.
            - The last time the user authenticated to any tailscale service, including the admin panel.
    currentlyConnected:
        type: boolean
        example: true
        description: |
            `true` when the user has a node currently connected to the control server.
```
