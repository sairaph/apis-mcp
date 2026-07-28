---
title: Get a user invite
page_id: operation-get-user-invites-userinviteid-7e8d9202
path: operations/userinvites
description: Retrieve a specific user invite.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /user-invites/{userInviteId}
operation_ids:
    - getUserInvite
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a user invite

`GET /user-invites/{userInviteId}`

Operation ID: `getUserInvite`

Retrieve a specific user invite.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userInviteId'
```

## Definition

```yaml
summary: Get a user invite
description: |
    Retrieve a specific user invite.
operationId: getUserInvite
tags:
    - UserInvites
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/UserInvite'
    '404':
        description: User invite not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
