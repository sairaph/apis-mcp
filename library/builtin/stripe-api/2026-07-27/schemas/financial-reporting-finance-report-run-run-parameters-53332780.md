---
title: financial_reporting_finance_report_run_run_parameters
page_id: schema-financial-reporting-finance-report-run-run-parameters-53332780
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# financial_reporting_finance_report_run_run_parameters

```yaml
{"title": "FinancialReportingFinanceReportRunRunParameters", "type": "object", "properties": {"columns": {"type": "array", "description": "The set of output columns requested for inclusion in the report run.", "items": {"maxLength": 5000, "type": "string"}}, "connected_account": {"maxLength": 5000, "type": "string", "description": "Connected account ID by which to filter the report run."}, "currency": {"type": "string", "description": "Currency of objects to be included in the report run.", "format": "currency"}, "interval_end": {"type": "integer", "description": "Ending timestamp of data to be included in the report run. Can be any UTC timestamp between 1 second after the user specified `interval_start` and 1 second before this report's last `data_available_end` value.", "format": "unix-time"}, "interval_start": {"type": "integer", "description": "Starting timestamp of data to be included in the report run. Can be any UTC timestamp between 1 second after this report's `data_available_start` and 1 second before the user specified `interval_end` value.", "format": "unix-time"}, "payout": {"maxLength": 5000, "type": "string", "description": "Payout ID by which to filter the report run."}, "reporting_category": {"maxLength": 5000, "type": "string", "description": "Category of balance transactions to be included in the report run."}, "timezone": {"maxLength": 5000, "type": "string", "description": "Defaults to `Etc/UTC`. The output timezone for all timestamps in the report. A list of possible time zone values is maintained at the [IANA Time Zone Database](http://www.iana.org/time-zones). Has no effect on `interval_start` or `interval_end`."}}, "description": "", "x-expandableFields": []}
```
