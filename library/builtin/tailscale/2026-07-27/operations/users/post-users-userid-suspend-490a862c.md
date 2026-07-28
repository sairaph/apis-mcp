---
title: Suspend a user
page_id: operation-post-users-userid-suspend-dbe33ece
path: operations/users
description: |-
    Suspends a user from their tailnet. Learn more about [suspending users](/kb/1145/remove-team-members#suspending-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot suspend their own user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /users/{userId}/suspend
operation_ids:
    - suspendUser
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Suspend a user

`POST /users/{userId}/suspend`

Operation ID: `suspendUser`

Suspends a user from their tailnet. Learn more about [suspending users](/kb/1145/remove-team-members#suspending-users).

OAuth Scope: `users`.

> ⓘ User-based access tokens cannot suspend their own user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Suspend a user
description: |
    Suspends a user from their tailnet. Learn more about [suspending users](/kb/1145/remove-team-members#suspending-users).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot suspend their own user.
operationId: suspendUser
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
