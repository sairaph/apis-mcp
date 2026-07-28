---
title: DebugEvent
page_id: schema-debugevent-9d950755
path: schemas
description: Debug event emitted when debug.echo_upstream_body is true. Contains the transformed upstream request body or timing milestones.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# DebugEvent

Debug event emitted when debug.echo_upstream_body is true. Contains the transformed upstream request body or timing milestones.

```yaml
{"description": "Debug event emitted when debug.echo_upstream_body is true. Contains the transformed upstream request body or timing milestones.", "example": {"debug": {"echo_upstream_body": {"messages": [], "model": "anthropic/claude-sonnet-4"}}, "sequence_number": 1, "type": "response.debug"}, "properties": {"debug": {"properties": {"echo_upstream_body": {"additionalProperties": {}, "type": "object"}, "timings": {"properties": {"epoch_ms": {"type": "integer"}, "event": {"enum": ["adapter_request", "upstream_headers_received", "first_token_received", "upstream_body_ended"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "start_ms": {"type": "integer"}}, "required": ["start_ms", "event", "epoch_ms"], "type": "object"}}, "type": "object"}, "sequence_number": {"type": "integer"}, "type": {"enum": ["response.debug"], "type": "string"}}, "required": ["type", "debug", "sequence_number"], "type": "object"}
```
