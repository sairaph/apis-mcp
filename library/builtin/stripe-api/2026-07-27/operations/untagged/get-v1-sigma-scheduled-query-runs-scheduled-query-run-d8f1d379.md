---
title: Retrieve a scheduled query run
page_id: operation-get-v1-sigma-scheduled-query-runs-scheduled-query-run-337791dc
path: operations/untagged
description: <p>Retrieves the details of an scheduled query run.</p>
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
http_methods:
    - GET
api_endpoints:
    - /v1/sigma/scheduled_query_runs/{scheduled_query_run}
operation_ids:
    - GetSigmaScheduledQueryRunsScheduledQueryRun
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# Retrieve a scheduled query run

`GET /v1/sigma/scheduled_query_runs/{scheduled_query_run}`

Operation ID: `GetSigmaScheduledQueryRunsScheduledQueryRun`

<p>Retrieves the details of an scheduled query run.</p>

## Definition

```yaml
{"summary": "Retrieve a scheduled query run", "description": "<p>Retrieves the details of an scheduled query run.</p>", "operationId": "GetSigmaScheduledQueryRunsScheduledQueryRun", "parameters": [{"name": "expand", "in": "query", "description": "Specifies which fields in the response should be expanded.", "required": false, "style": "deepObject", "explode": true, "schema": {"type": "array", "items": {"maxLength": 5000, "type": "string"}}}, {"name": "scheduled_query_run", "in": "path", "required": true, "style": "simple", "explode": false, "schema": {"maxLength": 5000, "type": "string"}}], "requestBody": {"content": {"application/x-www-form-urlencoded": {"schema": {"type": "object", "properties": {}, "additionalProperties": false}, "encoding": {}}}, "required": false}, "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/scheduled_query_run"}}}}, "default": {"description": "Error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/error"}}}}}}
```
