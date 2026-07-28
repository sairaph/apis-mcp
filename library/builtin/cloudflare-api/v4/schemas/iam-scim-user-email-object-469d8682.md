---
title: iam_scim_user_email_object
page_id: schema-iam-scim-user-email-object-469d8682
path: schemas
description: An email address entry for a SCIM User.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_user_email_object

An email address entry for a SCIM User.

```yaml
{"description": "An email address entry for a SCIM User.", "type": "object", "properties": {"primary": {"description": "A Boolean value indicating the preferred email address.", "type": "boolean", "example": true}, "type": {"description": "A label indicating the attribute's function, e.g., \"work\" or \"home\".", "type": "string", "example": "work"}, "value": {"description": "The email address value.", "type": "string", "format": "email", "example": "user@example.com"}}, "required": ["value"], "title": "SCIM User Email"}
```
