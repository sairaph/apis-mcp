---
title: rlanalytics_rate_limit_analytics_action_counters
page_id: schema-rlanalytics-rate-limit-analytics-action-counters-6fc9e3d9
path: schemas
description: Rate limiting action counts broken down by action type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rlanalytics_rate_limit_analytics_action_counters

Rate limiting action counts broken down by action type.

```yaml
{"description": "Rate limiting action counts broken down by action type.", "type": "object", "properties": {"all": {"description": "Count of rule actions of any type triggered.", "type": "integer", "format": "int64", "example": 700}, "allow": {"description": "Count of allow rule actions triggered.", "type": "integer", "format": "int64", "example": 500}, "ban": {"description": "Count of ban rule actions triggered.", "type": "integer", "format": "int64", "example": 50}, "error": {"description": "Count of error rule actions triggered.", "type": "integer", "format": "int64", "example": 50}, "simulate": {"description": "Count of simulate rule actions triggered.", "type": "integer", "format": "int64", "example": 100}}}
```
