---
title: workers_limits
page_id: schema-workers-limits-98eba5c8
path: schemas
description: Limits to apply for this Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_limits

Limits to apply for this Worker.

```yaml
{"description": "Limits to apply for this Worker.", "type": "object", "properties": {"cpu_ms": {"description": "The amount of CPU time this Worker can use in milliseconds.", "type": "integer", "example": 50, "x-auditable": true}, "subrequests": {"description": "The number of subrequests this Worker can make per request.", "type": "integer", "example": 1000, "x-auditable": true}}}
```
