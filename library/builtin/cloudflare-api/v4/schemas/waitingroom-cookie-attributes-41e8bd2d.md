---
title: waitingroom_cookie_attributes
page_id: schema-waitingroom-cookie-attributes-41e8bd2d
path: schemas
description: Configures cookie attributes for the waiting room cookie. This encrypted cookie stores a user's status in the waiting room, such as queue position.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waitingroom_cookie_attributes

Configures cookie attributes for the waiting room cookie. This encrypted cookie stores a user's status in the waiting room, such as queue position.

```yaml
{"description": "Configures cookie attributes for the waiting room cookie. This encrypted cookie stores a user's status in the waiting room, such as queue position.", "type": "object", "properties": {"samesite": {"description": "Configures the SameSite attribute on the waiting room cookie. Value `auto` will be translated to `lax` or `none` depending if **Always Use HTTPS** is enabled. Note that when using value `none`, the secure attribute cannot be set to `never`.", "type": "string", "example": "auto", "default": "auto", "enum": ["auto", "lax", "none", "strict"], "x-auditable": true}, "secure": {"description": "Configures the Secure attribute on the waiting room cookie. Value `always` indicates that the Secure attribute will be set in the Set-Cookie header, `never` indicates that the Secure attribute will not be set, and `auto` will set the Secure attribute depending if **Always Use HTTPS** is enabled.", "type": "string", "example": "auto", "default": "auto", "enum": ["auto", "always", "never"], "x-auditable": true}}}
```
