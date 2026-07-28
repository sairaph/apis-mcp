---
title: UsageMetrics
page_id: schema-usagemetrics-d5db18ad
path: schemas
source: https://docs.tavily.com/documentation/api-reference/openapi.json
source_type: openapi
imported_from: https://docs.tavily.com/documentation/api-reference/openapi.json
---

# UsageMetrics

```yaml
{"type": "object", "properties": {"usage": {"type": "integer", "description": "Credits consumed", "example": 543}, "paygo_cost_usd": {"type": "number", "description": "Pay-as-you-go cost in USD (overage credits billed at the account's per-credit rate)", "example": 4.34}, "request_count": {"type": "integer", "description": "Number of requests", "example": 528}}}
```
