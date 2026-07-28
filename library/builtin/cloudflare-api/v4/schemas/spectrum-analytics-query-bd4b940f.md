---
title: spectrum-analytics_query
page_id: schema-spectrum-analytics-query-bd4b940f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-analytics_query

```yaml
{"type": "object", "properties": {"dimensions": {"$ref": "#/components/schemas/spectrum-analytics_dimensions"}, "filters": {"$ref": "#/components/schemas/spectrum-analytics_filters"}, "limit": {"description": "Limit number of returned metrics.", "type": "number"}, "metrics": {"$ref": "#/components/schemas/spectrum-analytics_metrics"}, "since": {"$ref": "#/components/schemas/spectrum-analytics_since"}, "sort": {"$ref": "#/components/schemas/spectrum-analytics_sort"}, "until": {"$ref": "#/components/schemas/spectrum-analytics_until"}}}
```
