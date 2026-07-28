---
title: access_require
page_id: schema-access-require-a4d3c2bf
path: schemas
description: Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_require

Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.

```yaml
{"description": "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.", "type": "array", "items": {"$ref": "#/components/schemas/access_rule"}}
```
