---
title: access_email_list_rule
page_id: schema-access-email-list-rule-192e75d2
path: schemas
description: Matches an email address from a list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_email_list_rule

Matches an email address from a list.

```yaml
{"description": "Matches an email address from a list.", "type": "object", "properties": {"email_list": {"type": "object", "properties": {"id": {"description": "The ID of a previously created email list.", "type": "string", "example": "aa0a4aab-672b-4bdb-bc33-a59f1130a11f"}}, "required": ["id"]}}, "required": ["email_list"], "title": "Email list"}
```
