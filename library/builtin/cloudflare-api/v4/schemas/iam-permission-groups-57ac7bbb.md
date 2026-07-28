---
title: iam_permission_groups
page_id: schema-iam-permission-groups-57ac7bbb
path: schemas
description: A set of permission groups that are specified to the policy.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_permission_groups

A set of permission groups that are specified to the policy.

```yaml
{"description": "A set of permission groups that are specified to the policy.", "type": "array", "items": {"$ref": "#/components/schemas/iam_permission_group"}, "example": [{"id": "c8fed203ed3043cba015a93ad1616f1f", "meta": {"label": "load_balancer_admin", "scopes": "com.cloudflare.api.account"}, "name": "Zone Read"}, {"id": "82e64a83756745bbbb1c9c2701bf816b", "meta": {"label": "fbm_user", "scopes": "com.cloudflare.api.account"}, "name": "Magic Network Monitoring"}]}
```
