---
title: observatory_availabilities
page_id: schema-observatory-availabilities-b4a0c97c
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# observatory_availabilities

```yaml
{"type": "object", "properties": {"quota": {"type": "object", "properties": {"plan": {"description": "Cloudflare plan.", "type": "string", "example": "free", "x-auditable": true}, "quotasPerPlan": {"description": "The number of tests available per plan.", "type": "object", "properties": {"value": {"$ref": "#/components/schemas/observatory_plan-properties-info"}}}, "remainingSchedules": {"description": "The number of remaining schedules available.", "type": "number", "example": 1, "x-auditable": true}, "remainingTests": {"description": "The number of remaining tests available.", "type": "number", "example": 30, "x-auditable": true}, "scheduleQuotasPerPlan": {"description": "The number of schedules available per plan.", "type": "object", "properties": {"value": {"$ref": "#/components/schemas/observatory_plan-properties-info"}}}}}, "regions": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_labeled_region"}}, "regionsPerPlan": {"description": "Available regions.", "type": "object", "properties": {"business": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_labeled_region"}}, "enterprise": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_labeled_region"}}, "free": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_labeled_region"}}, "pro": {"type": "array", "items": {"$ref": "#/components/schemas/observatory_labeled_region"}}}}}}
```
