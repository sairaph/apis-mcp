---
title: zone-analytics-api_user_period_bandwidth
page_id: schema-zone-analytics-api-user-period-bandwidth-a933a7a4
path: schemas
description: Breakdown of bandwidth totals for a time period, in bytes.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zone-analytics-api_user_period_bandwidth

Breakdown of bandwidth totals for a time period, in bytes.

```yaml
{"description": "Breakdown of bandwidth totals for a time period, in bytes.", "type": "object", "properties": {"all": {"description": "The total number of bytes served.", "type": "integer", "example": 190290}, "cached": {"description": "The number of bytes served from cache.", "type": "integer", "example": 97717}, "uncached": {"description": "The number of bytes served from the origin.", "type": "integer", "example": 92573}}}
```
