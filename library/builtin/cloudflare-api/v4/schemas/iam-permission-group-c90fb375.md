---
title: iam_permission_group
page_id: schema-iam-permission-group-c90fb375
path: schemas
description: A named group of permissions that map to a group of operations against resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_permission_group

A named group of permissions that map to a group of operations against resources.

```yaml
{"description": "A named group of permissions that map to a group of operations against resources.", "type": "object", "properties": {"id": {"description": "Identifier of the permission group.", "type": "string", "example": "6d7f2f5f5b1d4a0e9081fdc98d432fd1", "x-auditable": true}, "meta": {"description": "Attributes associated to the permission group.", "type": "object", "example": {"label": "load_balancer_admin", "scopes": "com.cloudflare.api.account"}, "properties": {"key": {"type": "string", "x-auditable": true}, "value": {"type": "string", "x-auditable": true}}}, "name": {"description": "Name of the permission group.", "type": "string", "example": "Load Balancer", "readOnly": true, "x-auditable": true}}, "required": ["id"]}
```
