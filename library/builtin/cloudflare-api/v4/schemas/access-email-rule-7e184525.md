---
title: access_email_rule
page_id: schema-access-email-rule-7e184525
path: schemas
description: Matches a specific email.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_email_rule

Matches a specific email.

```yaml
{"description": "Matches a specific email.", "type": "object", "properties": {"email": {"type": "object", "properties": {"email": {"description": "The email of the user.", "type": "string", "format": "email", "example": "test@example.com"}}, "required": ["email"]}}, "required": ["email"], "title": "Email"}
```
