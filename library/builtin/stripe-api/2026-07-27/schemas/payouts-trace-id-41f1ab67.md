---
title: payouts_trace_id
page_id: schema-payouts-trace-id-41f1ab67
path: schemas
source: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/stripe/openapi/a2de9917ac9c7be3ba11abd5151b5c2df3add59e/latest/openapi.spec3.json
---

# payouts_trace_id

```yaml
{"title": "PayoutsTraceID", "required": ["status"], "type": "object", "properties": {"status": {"maxLength": 5000, "type": "string", "description": "Possible values are `pending`, `supported`, and `unsupported`. When `payout.status` is `pending` or `in_transit`, this will be `pending`. When the payout transitions to `paid`, `failed`, or `canceled`, this status will become `supported` or `unsupported` shortly after in most cases. In some cases, this may appear as `pending` for up to 10 days after `arrival_date` until transitioning to `supported` or `unsupported`."}, "value": {"maxLength": 5000, "type": "string", "description": "The trace ID value if `trace_id.status` is `supported`, otherwise `nil`.", "nullable": true}}, "description": "", "x-expandableFields": []}
```
