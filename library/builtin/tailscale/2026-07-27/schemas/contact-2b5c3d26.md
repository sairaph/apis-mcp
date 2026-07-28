---
title: Contact
page_id: schema-contact-2b5c3d26
path: schemas
description: A tailnet contact.
source: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
source_type: openapi
imported_from: https://api.tailscale.com/api/v2?outputOpenapiSchema=true
---

# Contact

A tailnet contact.

```yaml
type: object
description: A tailnet contact.
properties:
    email:
        type: string
        description: The contact's email address.
        example: user@example.com
    fallbackEmail:
        type: string
        description: The email address used when contact's email address has not been verified.
        example: otheruser@example.com
    needsVerification:
        type: boolean
        description: Indicates whether the contact's email address still needs to be verified.
        example: true
example:
    email: user@example.com
    fallbackEmail: otheruser@example.com"
    needsVerification: true
```
