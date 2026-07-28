---
title: Update contact
page_id: operation-patch-tailnet-tailnet-contacts-contacttype-f8030d76
path: operations/contacts
description: |-
    Update the preferences for this type of contact. If the email address has changed, the system will send a verification email to confirm the change.

    OAuth Scope: `account_settings`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - PATCH
api_endpoints:
    - /tailnet/{tailnet}/contacts/{contactType}
operation_ids:
    - updateContact
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Update contact

`PATCH /tailnet/{tailnet}/contacts/{contactType}`

Operation ID: `updateContact`

Update the preferences for this type of contact. If the email address has changed, the system will send a verification email to confirm the change.

OAuth Scope: `account_settings`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/contactType'
```

## Definition

```yaml
summary: Update contact
description: |
    Update the preferences for this type of contact. If the email address has changed, the system will send a verification email to confirm the change.

    OAuth Scope: `account_settings`.
operationId: updateContact
tags:
    - Contacts
requestBody:
    content:
        application/json:
            schema:
                type: object
                properties:
                    email:
                        type: string
                        description: The contact's email address.
                        example: newuser@example.com
                required:
                    - email
            example:
                email: newuser@example.com
responses:
    '200':
        description: Successful operation.
    '403':
        description: User does not have sufficient access to update contacts for this tailnet.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
