---
title: List users
page_id: operation-get-tailnet-tailnet-users-3dac8d08
path: operations/users
description: |-
    List all users of a tailnet.

    OAuth Scope: `users:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/users
operation_ids:
    - listUsers
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# List users

`GET /tailnet/{tailnet}/users`

Operation ID: `listUsers`

List all users of a tailnet.

OAuth Scope: `users:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- name: type
  in: query
  description: |
    Allows for filtering the output by user type.
  required: false
  schema:
    type: string
    enum:
        - member
        - shared
        - all
    default: member
    example: member
- name: role
  in: query
  description: |
    Allows for filtering the output by user role. Learn more about [user roles](kb/1138/user-roles).
  required: false
  schema:
    type: string
    enum:
        - owner
        - member
        - admin
        - it-admin
        - network-admin
        - billing-admin
        - auditor
        - all
    default: all
    example: all
```

## Definition

```yaml
summary: List users
description: |
    List all users of a tailnet.

    OAuth Scope: `users:read`.
operationId: listUsers
tags:
    - Users
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        users:
                            type: array
                            items:
                                $ref: '#/components/schemas/User'
    '400':
        $ref: '#/components/responses/400'
    '403':
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found, or user does not have access to read users.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
