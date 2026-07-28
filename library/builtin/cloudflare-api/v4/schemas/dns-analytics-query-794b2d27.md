---
title: dns-analytics_query
page_id: schema-dns-analytics-query-794b2d27
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-analytics_query

```yaml
{"type": "object", "properties": {"dimensions": {"description": "Array of dimension names.", "type": "array", "items": {"description": "Dimension name.", "example": "responseCode", "type": "string"}, "example": ["responseCode", "queryName"]}, "filters": {"$ref": "#/components/schemas/dns-analytics_filters"}, "limit": {"$ref": "#/components/schemas/dns-analytics_limit"}, "metrics": {"description": "Array of metric names.", "type": "array", "items": {"description": "Metric name.", "example": "queries", "type": "string"}, "example": ["queryCount", "responseTimeAvg"]}, "since": {"$ref": "#/components/schemas/dns-analytics_since"}, "sort": {"description": "Array of dimensions to sort by, where each dimension may be prefixed by - (descending) or + (ascending).", "type": "array", "items": {"description": "Dimension name (may be prefixed by - (descending) or + (ascending).", "example": "+responseCode", "type": "string"}, "example": ["+responseCode", "-queryName"]}, "until": {"$ref": "#/components/schemas/dns-analytics_until"}}, "required": ["dimensions", "metrics", "since", "until", "limit"]}
```
