---
title: reporting.report_type
page_id: schema-reporting-report-type-8c77bfee
path: schemas
description: |-
    The Report Type resource corresponds to a particular type of report, such as
    the "Activity summary" or "Itemized payouts" reports. These objects are
    identified by an ID belonging to a set of enumerated values. See
    [API Access to Reports documentation](https://docs.stripe.com/reporting/statements/api)
    for those Report Type IDs, along with required and optional parameters.

    Note that certain report types can only be run based on your live-mode data (not test-mode
    data), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# reporting.report_type

The Report Type resource corresponds to a particular type of report, such as
the "Activity summary" or "Itemized payouts" reports. These objects are
identified by an ID belonging to a set of enumerated values. See
[API Access to Reports documentation](https://docs.stripe.com/reporting/statements/api)
for those Report Type IDs, along with required and optional parameters.

Note that certain report types can only be run based on your live-mode data (not test-mode
data), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).

```yaml
{"title": "reporting_report_type", "required": ["data_available_end", "data_available_start", "id", "livemode", "name", "object", "updated", "version"], "type": "object", "properties": {"data_available_end": {"type": "integer", "description": "Most recent time for which this Report Type is available. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "data_available_start": {"type": "integer", "description": "Earliest time for which this Report Type is available. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "default_columns": {"type": "array", "description": "List of column names that are included by default when this Report Type gets run. (If the Report Type doesn't support the `columns` parameter, this will be null.)", "nullable": true, "items": {"maxLength": 5000, "type": "string"}}, "id": {"maxLength": 5000, "type": "string", "description": "The [ID of the Report Type](https://docs.stripe.com/reporting/statements/api#available-report-types), such as `balance.summary.1`."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "name": {"maxLength": 5000, "type": "string", "description": "Human-readable name of the Report Type"}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["reporting.report_type"]}, "updated": {"type": "integer", "description": "When this Report Type was latest updated. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "version": {"type": "integer", "description": "Version of the Report Type. Different versions report with the same ID will have the same purpose, but may take different run parameters or have different result schemas."}}, "description": "The Report Type resource corresponds to a particular type of report, such as\nthe \"Activity summary\" or \"Itemized payouts\" reports. These objects are\nidentified by an ID belonging to a set of enumerated values. See\n[API Access to Reports documentation](https://docs.stripe.com/reporting/statements/api)\nfor those Report Type IDs, along with required and optional parameters.\n\nNote that certain report types can only be run based on your live-mode data (not test-mode\ndata), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).", "x-expandableFields": [], "x-resourceId": "reporting.report_type"}
```
