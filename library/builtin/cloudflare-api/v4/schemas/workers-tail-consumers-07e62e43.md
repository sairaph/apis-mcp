---
title: workers_tail_consumers
page_id: schema-workers-tail-consumers-07e62e43
path: schemas
description: List of Workers that will consume logs from the attached Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_tail_consumers

List of Workers that will consume logs from the attached Worker.

```yaml
{"description": "List of Workers that will consume logs from the attached Worker.", "type": "array", "items": {"$ref": "#/components/schemas/workers_tail_consumers_script"}, "nullable": true, "x-stainless-collection-type": "set"}
```
