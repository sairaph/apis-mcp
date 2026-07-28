---
title: aaa_unsubscribe_email
page_id: schema-aaa-unsubscribe-email-7713452a
path: schemas
description: Response body for the GET show-unsubscribe-details endpoint. All fields are populated.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_unsubscribe_email

Response body for the GET show-unsubscribe-details endpoint. All fields are populated.

```yaml
{"description": "Response body for the GET show-unsubscribe-details endpoint. All fields are populated.", "type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/aaa_account-id"}, "email": {"type": "string", "format": "email"}, "id": {"$ref": "#/components/schemas/aaa_policy-id"}, "name": {"$ref": "#/components/schemas/aaa_schemas-name"}, "token": {"type": "string", "x-sensitive": true}}}
```
