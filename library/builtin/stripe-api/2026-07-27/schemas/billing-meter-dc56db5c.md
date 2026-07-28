---
title: billing.meter
page_id: schema-billing-meter-dc56db5c
path: schemas
description: |-
    Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.

    Related guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.meter

Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.

Related guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)

```yaml
{"title": "BillingMeter", "required": ["created", "customer_mapping", "default_aggregation", "display_name", "event_name", "id", "livemode", "object", "status", "status_transitions", "updated", "value_settings"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "customer_mapping": {"$ref": "#/components/schemas/billing_meter_resource_customer_mapping_settings"}, "default_aggregation": {"$ref": "#/components/schemas/billing_meter_resource_aggregation_settings"}, "display_name": {"maxLength": 5000, "type": "string", "description": "The meter's name."}, "event_name": {"maxLength": 5000, "type": "string", "description": "The name of the meter event to record usage for. Corresponds with the `event_name` field on meter events."}, "event_time_window": {"type": "string", "description": "The time window which meter events have been pre-aggregated for, if any.", "nullable": true, "enum": ["day", "hour"]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.meter"]}, "status": {"type": "string", "description": "The meter's status.", "enum": ["active", "inactive"]}, "status_transitions": {"$ref": "#/components/schemas/billing_meter_resource_billing_meter_status_transitions"}, "updated": {"type": "integer", "description": "Time at which the object was last updated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "value_settings": {"$ref": "#/components/schemas/billing_meter_resource_billing_meter_value"}}, "description": "Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.\n\nRelated guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)", "x-expandableFields": ["customer_mapping", "default_aggregation", "status_transitions", "value_settings"], "x-resourceId": "billing.meter"}
```
