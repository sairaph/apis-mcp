---
title: workers_observability
page_id: schema-workers-observability-1f8695bb
path: schemas
description: Observability settings for the Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_observability

Observability settings for the Worker.

```yaml
{"description": "Observability settings for the Worker.", "type": "object", "properties": {"enabled": {"description": "Whether observability is enabled for the Worker.", "type": "boolean", "example": true, "x-auditable": true}, "head_sampling_rate": {"description": "The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.", "type": "number", "example": 0.1, "nullable": true, "x-auditable": true}, "logs": {"description": "Log settings for the Worker.", "type": "object", "nullable": true, "properties": {"destinations": {"description": "A list of destinations where logs will be exported to.", "type": "array", "items": {"type": "string"}, "example": ["cloudflare"], "x-auditable": true}, "enabled": {"description": "Whether logs are enabled for the Worker.", "type": "boolean", "example": true, "x-auditable": true}, "head_sampling_rate": {"description": "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.", "type": "number", "example": 0.1, "nullable": true, "x-auditable": true}, "invocation_logs": {"description": "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.", "type": "boolean", "example": true, "x-auditable": true}, "persist": {"description": "Whether log persistence is enabled for the Worker.", "type": "boolean", "example": true, "default": true, "x-auditable": true}}, "required": ["enabled", "invocation_logs"]}, "traces": {"description": "Trace settings for the Worker.", "type": "object", "nullable": true, "properties": {"destinations": {"description": "A list of destinations where traces will be exported to.", "type": "array", "items": {"type": "string"}, "example": ["cloudflare"], "x-auditable": true}, "enabled": {"description": "Whether traces are enabled for the Worker.", "type": "boolean", "example": true, "x-auditable": true}, "head_sampling_rate": {"description": "The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.", "type": "number", "example": 0.1, "nullable": true, "x-auditable": true}, "persist": {"description": "Whether trace persistence is enabled for the Worker.", "type": "boolean", "example": true, "default": true, "x-auditable": true}, "propagation_policy": {"description": "Controls how inbound trace context (traceparent/tracestate) headers on incoming requests are handled. \"authenticated\" (default) honors inbound trace context only when accompanied by a valid trace auth token. \"accept\" unconditionally accepts inbound trace context. Requires the trace propagation feature to be enabled.", "type": "string", "default": "authenticated", "enum": ["authenticated", "accept"], "x-auditable": true}}}}, "required": ["enabled"]}
```
