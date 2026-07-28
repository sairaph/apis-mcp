---
title: bot-management_requests_by_score
page_id: schema-bot-management-requests-by-score-519b9d0b
path: schemas
description: Map of bot scores (1-99) to request counts. Sum must equal `requests`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# bot-management_requests_by_score

Map of bot scores (1-99) to request counts. Sum must equal `requests`.

```yaml
{"description": "Map of bot scores (1-99) to request counts. Sum must equal `requests`.", "type": "object", "additionalProperties": {"format": "int64", "type": "integer"}}
```
