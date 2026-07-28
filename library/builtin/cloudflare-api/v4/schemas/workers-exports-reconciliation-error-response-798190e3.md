---
title: workers_exports_reconciliation_error_response
page_id: schema-workers-exports-reconciliation-error-response-798190e3
path: schemas
description: |-
    The v4 error envelope returned with HTTP 400 when one or more
    declared `exports` entries fail reconciliation (error code
    100402). The top-level error `message` aggregates the per-class
    failures as human-readable prose; the structured per-class payload
    is carried in `errors[].meta.details`, sorted by class name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_exports_reconciliation_error_response

The v4 error envelope returned with HTTP 400 when one or more
declared `exports` entries fail reconciliation (error code
100402). The top-level error `message` aggregates the per-class
failures as human-readable prose; the structured per-class payload
is carried in `errors[].meta.details`, sorted by class name.

```yaml
{"description": "The v4 error envelope returned with HTTP 400 when one or more\ndeclared `exports` entries fail reconciliation (error code\n100402). The top-level error `message` aggregates the per-class\nfailures as human-readable prose; the structured per-class payload\nis carried in `errors[].meta.details`, sorted by class name.\n", "type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"code": {"type": "integer", "example": 100402, "enum": [100402]}, "message": {"description": "Aggregate, multi-line, human-readable summary of every\nper-class reconciliation failure.\n", "type": "string"}, "meta": {"type": "object", "properties": {"details": {"description": "Structured per-class reconciliation failures, sorted\nby class name.\n", "type": "array", "items": {"$ref": "#/components/schemas/workers_exports_reconciliation_error"}}}, "required": ["details"]}}, "required": ["code", "message", "meta"], "type": "object"}}, "messages": {"type": "array", "items": {}}, "result": {"type": "object", "enum": [null], "nullable": true}, "success": {"type": "boolean", "example": false, "enum": [false]}}, "required": ["success", "errors", "messages", "result"]}
```
