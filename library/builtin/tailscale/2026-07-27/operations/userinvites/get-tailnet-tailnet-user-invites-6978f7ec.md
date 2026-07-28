---
title: List user invites
page_id: operation-get-tailnet-tailnet-user-invites-8c7fbad0
path: operations/userinvites
description: List all open (not yet accepted) user invites to the tailnet.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/user-invites
operation_ids:
    - listUserInvites
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List user invites

`GET /tailnet/{tailnet}/user-invites`

Operation ID: `listUserInvites`

List all open (not yet accepted) user invites to the tailnet.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: List user invites
description: List all open (not yet accepted) user invites to the tailnet.
operationId: listUserInvites
tags:
    - UserInvites
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: array
                    items:
                        $ref: '#/components/schemas/UserInvite'
                    example:
                        - id: '29214'
                          role: admin
                          tailnetId: 12345
                          inviterId: 34567
                          email: user@example.com
                          lastEmailSentAt: '2024-05-09T16:23:26.91778771Z'
                          inviteUrl: https://login.tailscale.com/uinv/<code>
                        - id: '29215'
                          role: admin
                          tailnetId: 12345
                          inviterId: 34567
                          email: someoneelse@example.com
                          lastEmailSentAt: '2024-05-09T17:23:30.91778771Z'
                          inviteUrl: https://login.tailscale.com/uinv/<code>
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
