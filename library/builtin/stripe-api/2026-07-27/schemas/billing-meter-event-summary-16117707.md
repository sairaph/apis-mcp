---
title: billing.meter_event_summary
page_id: schema-billing-meter-event-summary-16117707
path: schemas
description: |-
    A billing meter event summary represents an aggregated view of a customer's billing meter events within a specified timeframe. It indicates how much
    usage was accrued by a customer for that period.

    Note: Meters events are aggregated asynchronously so the meter event summaries provide an eventually consistent view of the reported usage.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# billing.meter_event_summary

A billing meter event summary represents an aggregated view of a customer's billing meter events within a specified timeframe. It indicates how much
usage was accrued by a customer for that period.

Note: Meters events are aggregated asynchronously so the meter event summaries provide an eventually consistent view of the reported usage.

```yaml
{"title": "BillingMeterEventSummary", "required": ["aggregated_value", "end_time", "id", "livemode", "meter", "object", "start_time"], "type": "object", "properties": {"aggregated_value": {"type": "number", "description": "Aggregated value of all the events within `start_time` (inclusive) and `end_time` (inclusive). The aggregation strategy is defined on meter via `default_aggregation`."}, "end_time": {"type": "integer", "description": "End timestamp for this event summary (exclusive). Must be aligned with minute boundaries.", "format": "unix-time"}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "meter": {"maxLength": 5000, "type": "string", "description": "The meter associated with this event summary."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["billing.meter_event_summary"]}, "start_time": {"type": "integer", "description": "Start timestamp for this event summary (inclusive). Must be aligned with minute boundaries.", "format": "unix-time"}}, "description": "A billing meter event summary represents an aggregated view of a customer's billing meter events within a specified timeframe. It indicates how much\nusage was accrued by a customer for that period.\n\nNote: Meters events are aggregated asynchronously so the meter event summaries provide an eventually consistent view of the reported usage.", "x-expandableFields": [], "x-resourceId": "billing.meter_event_summary"}
```
