---
title: Update user role
page_id: operation-post-users-userid-role-0a6c05ae
path: operations/users
description: |-
    Update the role for the specified user.

    Learn more about [user roles](kb/1138/user-roles).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot update their own user's role.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /users/{userId}/role
operation_ids:
    - updateUserRole
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update user role

`POST /users/{userId}/role`

Operation ID: `updateUserRole`

Update the role for the specified user.

Learn more about [user roles](kb/1138/user-roles).

OAuth Scope: `users`.

> ⓘ User-based access tokens cannot update their own user's role.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Update user role
description: |
    Update the role for the specified user.

    Learn more about [user roles](kb/1138/user-roles).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot update their own user's role.
operationId: updateUserRole
tags:
    - Users
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
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
responses:
    '200':
        description: Successful operation.
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: User not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
