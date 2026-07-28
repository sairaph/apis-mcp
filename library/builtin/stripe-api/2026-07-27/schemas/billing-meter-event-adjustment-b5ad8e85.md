---
title: billing.meter_event_adjustment
page_id: schema-billing-meter-event-adjustment-b5ad8e85
path: schemas
description: A billing meter event adjustment is a resource that allows you to cancel a meter event. For example, you might create a billing meter event adjustment to cancel a meter event that was created in error or attached to the wrong customer.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.meter_event_adjustment

A billing meter event adjustment is a resource that allows you to cancel a meter event. For example, you might create a billing meter event adjustment to cancel a meter event that was created in error or attached to the wrong customer.

```yaml
{"title": "BillingMeterEventAdjustment", "required": ["event_name", "livemode", "object", "status", "type"], "type": "object", "properties": {"cancel": {"description": "Specifies which event to cancel.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/billing_meter_resource_billing_meter_event_adjustment_cancel"}]}, "event_name": {"maxLength": 100, "type": "string", "description": "The name of the meter event. Corresponds with the `event_name` field on a meter."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.meter_event_adjustment"]}, "status": {"type": "string", "description": "The meter event adjustment's status.", "enum": ["complete", "pending"]}, "type": {"type": "string", "description": "Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.", "enum": ["cancel"]}}, "description": "A billing meter event adjustment is a resource that allows you to cancel a meter event. For example, you might create a billing meter event adjustment to cancel a meter event that was created in error or attached to the wrong customer.", "x-expandableFields": ["cancel"], "x-resourceId": "billing.meter_event_adjustment"}
```
