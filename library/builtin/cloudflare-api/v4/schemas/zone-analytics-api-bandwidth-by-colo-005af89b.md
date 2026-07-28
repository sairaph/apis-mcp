---
title: zone-analytics-api_bandwidth_by_colo
page_id: schema-zone-analytics-api-bandwidth-by-colo-005af89b
path: schemas
description: Breakdown of totals for bandwidth in the form of bytes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_bandwidth_by_colo

Breakdown of totals for bandwidth in the form of bytes.

```yaml
{"description": "Breakdown of totals for bandwidth in the form of bytes.", "type": "object", "properties": {"all": {"description": "The total number of bytes served within the time frame.", "type": "integer"}, "cached": {"description": "The number of bytes that were cached (and served) by Cloudflare.", "type": "integer"}, "uncached": {"description": "The number of bytes that were fetched and served from the origin server.", "type": "integer"}}}
```
