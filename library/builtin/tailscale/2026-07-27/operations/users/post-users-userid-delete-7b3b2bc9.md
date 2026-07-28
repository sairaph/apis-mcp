---
title: Delete a user
page_id: operation-post-users-userid-delete-60f3a699
path: operations/users
description: |-
    Delete a user from their tailnet. Learn more about [deleting users](/kb/1145/remove-team-members#deleting-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot delete their own user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /users/{userId}/delete
operation_ids:
    - deleteUser
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Delete a user

`POST /users/{userId}/delete`

Operation ID: `deleteUser`

Delete a user from their tailnet. Learn more about [deleting users](/kb/1145/remove-team-members#deleting-users).

OAuth Scope: `users`.

> ⓘ User-based access tokens cannot delete their own user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Delete a user
description: |
    Delete a user from their tailnet. Learn more about [deleting users](/kb/1145/remove-team-members#deleting-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot delete their own user.
operationId: deleteUser
tags:
    - Users
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
