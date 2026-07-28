---
title: art_SummaryResult
page_id: schema-art-summaryresult-2653220c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# art_SummaryResult

```yaml
{"type": "object", "properties": {"currentTotal": {"description": "Aggregated stats for the requested time range.", "type": "array", "items": {"$ref": "#/components/schemas/art_ResultValues"}, "example": [{"attemptsTotal": 48291}]}, "previousTotal": {"description": "Aggregated stats for the equivalent preceding time range, for trend comparison.\n", "type": "array", "items": {"$ref": "#/components/schemas/art_ResultValues"}, "example": [{"attemptsTotal": 41033}]}}, "required": ["currentTotal", "previousTotal"]}
```
