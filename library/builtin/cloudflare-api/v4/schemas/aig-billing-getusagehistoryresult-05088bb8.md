---
title: aig-billing_GetUsageHistoryResult
page_id: schema-aig-billing-getusagehistoryresult-05088bb8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aig-billing_GetUsageHistoryResult

```yaml
{"type": "object", "properties": {"history": {"type": "array", "items": {"properties": {"aggregated_value": {"type": "number"}, "end_time": {"type": "number"}, "id": {"type": "string"}, "start_time": {"type": "number"}}, "required": ["id", "aggregated_value", "end_time", "start_time"], "type": "object"}}}, "required": ["history"]}
```
