---
title: Retrieve a Report Run
page_id: operation-get-v1-reporting-report-runs-report-run-26c26a0f
path: operations/untagged
description: <p>Retrieves the details of an existing Report Run.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/reporting/report_runs/{report_run}
operation_ids:
    - GetReportingReportRunsReportRun
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a Report Run

`GET /v1/reporting/report_runs/{report_run}`

Operation ID: `GetReportingReportRunsReportRun`

<p>Retrieves the details of an existing Report Run.</p>

## Definition

```yaml
{"summary": "Retrieve a Report Run", "description": "<p>Retrieves the details of an existing Report Run.</p>", "operationId": "GetReportingReportRunsReportRun", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "report_run", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/reporting.report_run"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
