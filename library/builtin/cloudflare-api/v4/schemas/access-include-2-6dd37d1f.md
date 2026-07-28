---
title: access_include-2
page_id: schema-access-include-2-6dd37d1f
path: schemas
description: Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_include-2

Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.

```yaml
{"description": "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.", "type": "array", "items": {"$ref": "#/components/schemas/access_rule"}, "default": [], "x-stainless-collection-type": "set"}
```
