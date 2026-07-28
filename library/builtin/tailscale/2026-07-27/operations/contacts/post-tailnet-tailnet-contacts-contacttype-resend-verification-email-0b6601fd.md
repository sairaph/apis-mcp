---
title: Resend verification email
page_id: operation-post-tailnet-tailnet-contacts-contacttype-resend-verification-email-033f2980
path: operations/contacts
description: |-
    Resends the verification email for this contact, if and only if verification is still pending.

    OAuth Scope: `account_settings`.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
http_methods:
    - POST
api_endpoints:
    - /tailnet/{tailnet}/contacts/{contactType}/resend-verification-email
operation_ids:
    - resendContactVerificationEmail
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Resend verification email

`POST /tailnet/{tailnet}/contacts/{contactType}/resend-verification-email`

Operation ID: `resendContactVerificationEmail`

Resends the verification email for this contact, if and only if verification is still pending.

OAuth Scope: `account_settings`.

## Path Parameters

```yaml
- $ref: '#/components/parameters/tailnet'
- $ref: '#/components/parameters/contactType'
```

## Definition

```yaml
summary: Resend verification email
description: |
    Resends the verification email for this contact, if and only if verification is still pending.

    OAuth Scope: `account_settings`.
operationId: resendContactVerificationEmail
tags:
    - Contacts
responses:
    '200':
        description: Successful operation.
    '400':
        description: Verification is not required, can't resend email.
        $ref: '#/components/responses/400'
    '403':
        description: User does not have sufficient access to update contacts for this tailnet.
        $ref: '#/components/responses/403'
    '404':
        description: Tailnet not found.
        $ref: '#/components/responses/404'
    '500':
        $ref: '#/components/responses/500'
```
