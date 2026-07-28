---
title: access_authentication_method_rule
page_id: schema-access-authentication-method-rule-2713155f
path: schemas
description: Enforce different MFA options
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_authentication_method_rule

Enforce different MFA options

```yaml
{"description": "Enforce different MFA options", "type": "object", "properties": {"auth_method": {"type": "object", "properties": {"auth_method": {"description": "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.", "type": "string", "example": "mfa"}}, "required": ["auth_method"]}}, "required": ["auth_method"], "title": "Authentication method"}
```
