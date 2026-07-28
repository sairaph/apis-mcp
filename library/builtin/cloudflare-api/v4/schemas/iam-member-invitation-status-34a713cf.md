---
title: iam_member-invitation-status
page_id: schema-iam-member-invitation-status-34a713cf
path: schemas
description: |-
    Status of the member invitation. If not provided during creation, defaults to 'pending'.
    Changing from 'accepted' back to 'pending' will trigger a replacement of the member resource in Terraform.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_member-invitation-status

Status of the member invitation. If not provided during creation, defaults to 'pending'.
Changing from 'accepted' back to 'pending' will trigger a replacement of the member resource in Terraform.

```yaml
{"description": "Status of the member invitation. If not provided during creation, defaults to 'pending'.\nChanging from 'accepted' back to 'pending' will trigger a replacement of the member resource in Terraform.\n", "type": "string", "enum": ["accepted", "pending"], "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}
```
