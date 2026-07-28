---
title: scheduled_query_run
page_id: schema-scheduled-query-run-35ea6808
path: schemas
description: |-
    If you have [scheduled a Sigma query](https://docs.stripe.com/sigma/scheduled-queries), you'll
    receive a `sigma.scheduled_query_run.created` webhook each time the query
    runs. The webhook contains a `ScheduledQueryRun` object, which you can use to
    retrieve the query results.
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# scheduled_query_run

If you have [scheduled a Sigma query](https://docs.stripe.com/sigma/scheduled-queries), you'll
receive a `sigma.scheduled_query_run.created` webhook each time the query
runs. The webhook contains a `ScheduledQueryRun` object, which you can use to
retrieve the query results.

```yaml
{"title": "ScheduledQueryRun", "required": ["created", "data_load_time", "id", "livemode", "object", "result_available_until", "sql", "status", "title"], "type": "object", "properties": {"created": {"type": "integer", "description": "Time at which the object was created. Measured in seconds since the Unix epoch.", "format": "unix-time"}, "data_load_time": {"type": "integer", "description": "When the query was run, Sigma contained a snapshot of your Stripe data at this time.", "format": "unix-time"}, "error": {"$ref": "#/components/schemas/sigma_scheduled_query_run_error"}, "file": {"description": "The file object representing the results of the query.", "nullable": true, "anyOf": [{"$ref": "#/components/schemas/file"}]}, "id": {"maxLength": 5000, "type": "string", "description": "Unique identifier for the object."}, "livemode": {"type": "boolean", "description": "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`."}, "object": {"type": "string", "description": "String representing the object's type. Objects of the same type share the same value.", "enum": ["scheduled_query_run"]}, "result_available_until": {"type": "integer", "description": "Time at which the result expires and is no longer available for download.", "format": "unix-time"}, "sql": {"maxLength": 100000, "type": "string", "description": "SQL for the query."}, "status": {"maxLength": 5000, "type": "string", "description": "The query's execution status, which will be `completed` for successful runs, and `canceled`, `failed`, or `timed_out` otherwise."}, "title": {"maxLength": 5000, "type": "string", "description": "Title of the query."}}, "description": "If you have [scheduled a Sigma query](https://docs.stripe.com/sigma/scheduled-queries), you'll\nreceive a `sigma.scheduled_query_run.created` webhook each time the query\nruns. The webhook contains a `ScheduledQueryRun` object, which you can use to\nretrieve the query results.", "x-expandableFields": ["error", "file"], "x-resourceId": "scheduled_query_run"}
```
