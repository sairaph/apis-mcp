---
title: recurring
page_id: schema-recurring-a4fdbb29
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# recurring

```yaml
{"title": "Recurring", "required": ["interval", "interval_count", "usage_type"], "type": "object", "properties": {"interval": {"type": "string", "description": "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.", "enum": ["day", "month", "week", "year"]}, "interval_count": {"type": "integer", "description": "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months."}, "meter": {"maxLength": 5000, "type": "string", "description": "The meter tracking the usage of a metered price", "nullable": true}, "usage_type": {"type": "string", "description": "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.", "enum": ["licensed", "metered"]}}, "description": "", "x-expandableFields": []}
```
