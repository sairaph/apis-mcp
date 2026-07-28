---
title: access_approval_groups-2
page_id: schema-access-approval-groups-2-ef48b682
path: schemas
description: Administrators who can approve a temporary authentication request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_approval_groups-2

Administrators who can approve a temporary authentication request.

```yaml
{"description": "Administrators who can approve a temporary authentication request.", "type": "array", "items": {"$ref": "#/components/schemas/access_approval_group-2"}, "example": [{"approvals_needed": 1, "email_addresses": ["test1@cloudflare.com", "test2@cloudflare.com"]}, {"approvals_needed": 3, "email_list_uuid": "597147a1-976b-4ef2-9af0-81d5d007fc34"}]}
```
