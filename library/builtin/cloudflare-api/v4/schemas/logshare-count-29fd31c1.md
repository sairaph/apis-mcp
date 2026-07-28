---
title: logshare_count
page_id: schema-logshare-count-29fd31c1
path: schemas
description: When `?count=` is provided, the response will contain up to `count` results. Since results are not sorted, you are likely to get different data for repeated requests. `count` must be an integer > 0.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logshare_count

When `?count=` is provided, the response will contain up to `count` results. Since results are not sorted, you are likely to get different data for repeated requests. `count` must be an integer > 0.

```yaml
{"description": "When `?count=` is provided, the response will contain up to `count` results. Since results are not sorted, you are likely to get different data for repeated requests. `count` must be an integer > 0.", "type": "integer", "minimum": 1, "x-auditable": true}
```
