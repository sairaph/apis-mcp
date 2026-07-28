---
title: reporting.report_run
page_id: schema-reporting-report-run-998ab6af
path: schemas
description: |-
    The Report Run object represents an instance of a report type generated with
    specific run parameters. Once the object is created, Stripe begins processing the report.
    When the report has finished running, it will give you a reference to a file
    where you can retrieve your results. For an overview, see
    [API Access to Reports](https://docs.stripe.com/reporting/statements/api).

    Note that certain report types can only be run based on your live-mode data (not test-mode
    data), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# reporting.report_run

The Report Run object represents an instance of a report type generated with
specific run parameters. Once the object is created, Stripe begins processing the report.
When the report has finished running, it will give you a reference to a file
where you can retrieve your results. For an overview, see
[API Access to Reports](https://docs.stripe.com/reporting/statements/api).

Note that certain report types can only be run based on your live-mode data (not test-mode
data), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).

```yaml
{"title": "reporting_report_run", "required": ["created", "id", "livemode", "object", "parameters", "report_type", "status"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "error": {"maxLength": 5000, "type": "string", "description": "If something should go wrong during the run, a message about the failure (populated when\n `status=failed`).", "nullable": true}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "`true` if the report is run on live mode data and `false` if it is run on test mode data."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["reporting.report_run"]}, "parameters": {"$ref": "#/components/schemas/financial_reporting_finance_report_run_run_parameters"}, "report_type": {"maxLength": 5000, "type": "string", "description": "The ID of the [report type](https://docs.stripe.com/reports/report-types) to run, such as `\"balance.summary.1\"`."}, "result": {"description": "The file object representing the result of the report run (populated when\n `status=succeeded`).", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/file"}]}, "status": {"maxLength": 5000, "type": "string", "description": "Status of this report run. This will be `pending` when the run is initially created.\n When the run finishes, this will be set to `succeeded` and the `result` field will be populated.\n Rarely, we may encounter an error, at which point this will be set to `failed` and the `error` field will be populated."}, "succeeded_at": {"type": "integer", "description": "Timestamp at which this run successfully finished (populated when\n `status=succeeded`). Measured in seconds since the Unix epoch.", "format": "unix-time", "nullable": true}}, "description": "The Report Run object represents an instance of a report type generated with\nspecific run parameters. Once the object is created, Stripe begins processing the report.\nWhen the report has finished running, it will give you a reference to a file\nwhere you can retrieve your results. For an overview, see\n[API Access to Reports](https://docs.stripe.com/reporting/statements/api).\n\nNote that certain report types can only be run based on your live-mode data (not test-mode\ndata), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).", "x-expandableFields": ["parameters", "result"], "x-resourceId": "reporting.report_run"}
```
