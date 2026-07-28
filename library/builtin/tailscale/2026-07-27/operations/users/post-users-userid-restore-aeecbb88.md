---
title: Restore a user
page_id: operation-post-users-userid-restore-be7b31d8
path: operations/users
description: |-
    Restores a suspended user's access to their tailnet. Learn more about [restoring users](/kb/1145/remove-team-members#restoring-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot restore their own user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /users/{userId}/restore
operation_ids:
    - restoreUser
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Restore a user

`POST /users/{userId}/restore`

Operation ID: `restoreUser`

Restores a suspended user's access to their tailnet. Learn more about [restoring users](/kb/1145/remove-team-members#restoring-users).

OAuth Scope: `users`.

> ⓘ User-based access tokens cannot restore their own user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Restore a user
description: |
    Restores a suspended user's access to their tailnet. Learn more about [restoring users](/kb/1145/remove-team-members#restoring-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot restore their own user.
operationId: restoreUser
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
