---
title: billing_meter_resource_aggregation_settings
page_id: schema-billing-meter-resource-aggregation-settings-3b0ab842
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing_meter_resource_aggregation_settings

```yaml
{"title": "BillingMeterResourceAggregationSettings", "required": ["formula"], "type": "object", "properties": {"formula": {"type": "string", "description": "Specifies how events are aggregated.", "enum": ["count", "last", "sum"], "x-stripeBypassValidation": true}}, "description": "", "x-expandableFields": []}
```
