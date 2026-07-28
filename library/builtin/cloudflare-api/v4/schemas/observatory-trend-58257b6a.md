---
title: observatory_trend
page_id: schema-observatory-trend-58257b6a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_trend

```yaml
{"type": "object", "properties": {"cls": {"description": "Cumulative Layout Shift trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "fcp": {"description": "First Contentful Paint trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "lcp": {"description": "Largest Contentful Paint trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "performanceScore": {"description": "The Lighthouse score trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "si": {"description": "Speed Index trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "tbt": {"description": "Total Blocking Time trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "ttfb": {"description": "Time To First Byte trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}, "tti": {"description": "Time To Interactive trend.", "type": "array", "items": {"nullable": true, "type": "number", "x-auditable": true}}}}
```
