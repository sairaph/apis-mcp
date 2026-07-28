---
title: workers_tags
page_id: schema-workers-tags-e0a23a12
path: schemas
description: Tags associated with the Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_tags

Tags associated with the Worker.

```yaml
{"description": "Tags associated with the Worker.", "type": "array", "items": {"$ref": "#/components/schemas/workers_tag"}, "example": ["my-team", "my-public-api"], "maxItems": 10, "nullable": true, "x-auditable": true, "x-stainless-collection-type": "set"}
```
