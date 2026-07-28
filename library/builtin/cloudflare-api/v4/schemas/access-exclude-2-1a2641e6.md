---
title: access_exclude-2
page_id: schema-access-exclude-2-1a2641e6
path: schemas
description: Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_exclude-2

Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.

```yaml
{"description": "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.", "type": "array", "items": {"$ref": "#/components/schemas/access_rule"}, "default": [], "x-stainless-collection-type": "set"}
```
