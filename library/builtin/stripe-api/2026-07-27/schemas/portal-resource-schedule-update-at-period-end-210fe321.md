---
title: portal_resource_schedule_update_at_period_end
page_id: schema-portal-resource-schedule-update-at-period-end-210fe321
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# portal_resource_schedule_update_at_period_end

```yaml
{"title": "PortalResourceScheduleUpdateAtPeriodEnd", "required": ["conditions"], "type": "object", "properties": {"conditions": {"type": "array", "description": "List of conditions. When any condition is true, an update will be scheduled at the end of the current period.", "items": {"$ref": "#/components/schemas/portal_resource_schedule_update_at_period_end_condition"}}}, "description": "", "x-expandableFields": ["conditions"]}
```
