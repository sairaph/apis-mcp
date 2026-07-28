---
title: thresholds_resource_usage_alert_filter
page_id: schema-thresholds-resource-usage-alert-filter-7972519a
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# thresholds_resource_usage_alert_filter

```yaml
{"title": "ThresholdsResourceUsageAlertFilter", "required": ["type"], "type": "object", "properties": {"customer": {"description": "Limit the scope of the alert to this customer ID", "nullable": true, "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/customer"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/customer"}]}}, "type": {"type": "string", "enum": ["customer"]}}, "description": "", "x-expandableFields": ["customer"]}
```
