---
title: Create user invites
page_id: operation-post-tailnet-tailnet-user-invites-1a5ad360
path: operations/userinvites
description: |-
    Create, and optionally email out, new user invites to join the tailnet.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/user-invites
operation_ids:
    - createUserInvites
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Create user invites

`POST /tailnet/{tailnet}/user-invites`

Operation ID: `createUserInvites`

Create, and optionally email out, new user invites to join the tailnet.

> ⓘ Only permitted for user-owned keys, because invites require an inviting user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Create user invites
description: |
    Create, and optionally email out, new user invites to join the tailnet.

    > ⓘ Only permitted for user-owned keys, because invites require an inviting user.
operationId: createUserInvites
tags:
    - UserInvites
requestBody:
    content:
        application/json:
            schema:
                type: array
                items:
                    type: object
                    properties:
                        role:
                            type: string
                            enum:
                                - member
                                - admin
                                - it-admin
                                - network-admin
                                - billing-admin
                                - auditor
                            default: member
                            example: admin
                            description: |
                                Optionally specifies a user role to assign the invited user.
                        email:
                            type: string
                            example: user@example.com
                            description: |
                                Optionally specifies the email to send the created invite.
                                If not set, the endpoint generates and returns an invite URL, but does not email it out.
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
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
