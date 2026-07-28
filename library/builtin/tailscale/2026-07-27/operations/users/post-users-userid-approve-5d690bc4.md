---
title: Approve a user
page_id: operation-post-users-userid-approve-9cbd3a26
path: operations/users
description: |-
    Approve a pending user's access to the tailnet.
    This is a no-op if user approval has not been enabled for the tailnet, or if the user is already approved.

    User approval can be managed using the [tailnet settings endpoints](#tag/tailnetsettings).

    Learn more about [user approval](/kb/1239/user-approval) and [enabling user approval for your network](/kb/1239/user-approval#enable-user-approval-for-your-network).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot approve their own user.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /users/{userId}/approve
operation_ids:
    - approveUser
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Approve a user

`POST /users/{userId}/approve`

Operation ID: `approveUser`

Approve a pending user's access to the tailnet.
This is a no-op if user approval has not been enabled for the tailnet, or if the user is already approved.

User approval can be managed using the [tailnet settings endpoints](#tag/tailnetsettings).

Learn more about [user approval](/kb/1239/user-approval) and [enabling user approval for your network](/kb/1239/user-approval#enable-user-approval-for-your-network).

OAuth Scope: `users`.

> ⓘ User-based access tokens cannot approve their own user.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Approve a user
description: |
    Approve a pending user's access to the tailnet.
    This is a no-op if user approval has not been enabled for the tailnet, or if the user is already approved.

    User approval can be managed using the [tailnet settings endpoints](#tag/tailnetsettings).

    Learn more about [user approval](/kb/1239/user-approval) and [enabling user approval for your network](/kb/1239/user-approval#enable-user-approval-for-your-network).

    OAuth Scope: `users`.

    > ⓘ User-based access tokens cannot approve their own user.
operationId: approveUser
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
