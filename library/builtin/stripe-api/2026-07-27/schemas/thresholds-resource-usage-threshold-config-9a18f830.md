---
title: thresholds_resource_usage_threshold_config
page_id: schema-thresholds-resource-usage-threshold-config-9a18f830
path: schemas
description: The usage threshold alert configuration enables setting up alerts for when a certain usage threshold on a specific meter is crossed.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# thresholds_resource_usage_threshold_config

The usage threshold alert configuration enables setting up alerts for when a certain usage threshold on a specific meter is crossed.

```yaml
{"title": "ThresholdsResourceUsageThresholdConfig", "required": ["gte", "meter", "recurrence"], "type": "object", "properties": {"filters": {"type": "array", "description": "The filters allow limiting the scope of this usage alert. You can only specify up to one filter at this time.", "nullable": true, "items": {"$ref": "#/components/schemas/thresholds_resource_usage_alert_filter"}}, "gte": {"type": "integer", "description": "The value at which this alert will trigger."}, "meter": {"description": "The [Billing Meter](/api/billing/meter) ID whose usage is monitored.", "anyOf": [{"maxLength": 5000, "type": "string"}, {"$ref": "#/components/schemas/billing.meter"}], "x-expansionResources": {"oneOf": [{"$ref": "#/components/schemas/billing.meter"}]}}, "recurrence": {"type": "string", "description": "Defines how the alert will behave.", "enum": ["one_time"], "x-stripeBypassValidation": true}}, "description": "The usage threshold alert configuration enables setting up alerts for when a certain usage threshold on a specific meter is crossed.", "x-expandableFields": ["filters", "meter"]}
```
