---
title: Get a user
page_id: operation-get-users-userid-e827a925
path: operations/users
description: |-
    Retrieve details about the specified user.

    OAuth Scope: `users:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /users/{userId}
operation_ids:
    - getUser
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get a user

`GET /users/{userId}`

Operation ID: `getUser`

Retrieve details about the specified user.

OAuth Scope: `users:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/userId'
```

## Definition

```yaml
summary: Get a user
description: |
    Retrieve details about the specified user.

    OAuth Scope: `users:read`.
operationId: getUser
tags:
    - Users
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    $ref: '#/components/schemas/User'
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
