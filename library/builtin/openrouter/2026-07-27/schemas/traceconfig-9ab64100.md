---
title: TraceConfig
page_id: schema-traceconfig-9ab64100
path: schemas
description: Metadata for observability and tracing. Known keys (trace_id, trace_name, span_name, generation_name, parent_span_id) have special handling. Additional keys are passed through as custom metadata to configured broadcast destinations.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# TraceConfig

Metadata for observability and tracing. Known keys (trace_id, trace_name, span_name, generation_name, parent_span_id) have special handling. Additional keys are passed through as custom metadata to configured broadcast destinations.

```yaml
{"additionalProperties": {}, "description": "Metadata for observability and tracing. Known keys (trace_id, trace_name, span_name, generation_name, parent_span_id) have special handling. Additional keys are passed through as custom metadata to configured broadcast destinations.", "example": {"trace_id": "trace-abc123", "trace_name": "my-app-trace"}, "properties": {"generation_name": {"type": "string"}, "parent_span_id": {"type": "string"}, "span_name": {"type": "string"}, "trace_id": {"type": "string"}, "trace_name": {"type": "string"}}, "type": "object"}
```
