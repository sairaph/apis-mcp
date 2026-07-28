---
title: access_approval_group-2
page_id: schema-access-approval-group-2-e28654fd
path: schemas
description: A group of email addresses that can approve a temporary authentication request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_approval_group-2

A group of email addresses that can approve a temporary authentication request.

```yaml
{"description": "A group of email addresses that can approve a temporary authentication request.", "type": "object", "properties": {"approvals_needed": {"description": "The number of approvals needed to obtain access.", "type": "number", "example": 1, "minimum": 0}, "email_addresses": {"description": "A list of emails that can approve the access request.", "type": "array", "items": {}, "example": ["test@cloudflare.com", "test2@cloudflare.com"]}, "email_list_uuid": {"description": "The UUID of an re-usable email list.", "type": "string"}}, "required": ["approvals_needed"]}
```
