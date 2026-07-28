---
title: Get contacts
page_id: operation-get-tailnet-tailnet-contacts-1c74155c
path: operations/contacts
description: |-
    Retrieve the tailnet's current contacts.

    OAuth Scope: `account_settings:read`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - GET
api_endpoints:
    - /tailnet/{tailnet}/contacts
operation_ids:
    - getContacts
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Get contacts

`GET /tailnet/{tailnet}/contacts`

Operation ID: `getContacts`

Retrieve the tailnet's current contacts.

OAuth Scope: `account_settings:read`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
```

## Definition

```yaml
summary: Get contacts
description: |
    Retrieve the tailnet's current contacts.

    OAuth Scope: `account_settings:read`.
operationId: getContacts
tags:
    - Contacts
responses:
    '200':
        description: Successful operation.
        content:
            application/json:
                schema:
                    type: object
                    properties:
                        account:
                            $ref: '#/components/schemas/Contact'
                        support:
                            $ref: '#/components/schemas/Contact'
                        security:
                            $ref: '#/components/schemas/Contact'
                example:
                    account:
                        email: owner@example.com
                    support:
                        email: support@example.com
                    security:
                        email: security@example.com
    '403':
        description: User does not have sufficient access to view contacts on this tailnet.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
