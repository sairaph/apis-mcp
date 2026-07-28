---
title: iam_scim_user_name_object
page_id: schema-iam-scim-user-name-object-e6f89c58
path: schemas
description: The components of the user's real name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_scim_user_name_object

The components of the user's real name.

```yaml
{"description": "The components of the user's real name.", "type": "object", "properties": {"familyName": {"description": "The family name (last name) of the user.", "type": "string", "example": "Smith", "x-auditable": true}, "formatted": {"description": "The full name, including all middle names, titles, and suffixes as appropriate, formatted for display.", "type": "string", "example": "Jane Smith", "readOnly": true}, "givenName": {"description": "The given name (first name) of the user.", "type": "string", "example": "Jane", "x-auditable": true}}, "title": "SCIM User Name"}
```
