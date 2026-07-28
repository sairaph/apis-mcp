---
title: workers_compatibility_flags
page_id: schema-workers-compatibility-flags-7e07965c
path: schemas
description: Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_compatibility_flags

Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.

```yaml
{"description": "Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.", "type": "array", "items": {"$ref": "#/components/schemas/workers_compatibility_flag"}, "example": ["nodejs_compat"], "default": [], "x-stainless-collection-type": "set"}
```
