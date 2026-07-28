---
title: access_resource_user_email
page_id: schema-access-resource-user-email-96c9babd
path: schemas
description: |-
    The email address of the SCIM User resource. Pass once for a single
    lookup (`?resource_user_email=A`) or repeat the parameter
    (`?resource_user_email=A&resource_user_email=B`) to filter by multiple
    emails in one request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_resource_user_email

The email address of the SCIM User resource. Pass once for a single
lookup (`?resource_user_email=A`) or repeat the parameter
(`?resource_user_email=A&resource_user_email=B`) to filter by multiple
emails in one request.

```yaml
{"description": "The email address of the SCIM User resource. Pass once for a single\nlookup (`?resource_user_email=A`) or repeat the parameter\n(`?resource_user_email=A&resource_user_email=B`) to filter by multiple\nemails in one request.", "type": "array", "items": {"format": "email", "type": "string"}, "example": ["john.smith@example.com"]}
```
