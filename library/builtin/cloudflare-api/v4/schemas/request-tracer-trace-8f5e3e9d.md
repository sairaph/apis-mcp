---
title: request-tracer_trace
page_id: schema-request-tracer-trace-8f5e3e9d
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# request-tracer_trace

```yaml
{"type": "array", "items": {"description": "List of steps acting on request/response", "properties": {"action": {"description": "If step type is rule, then action performed by this rule", "type": "string", "example": "execute", "pattern": "^[a-z_]+$", "x-auditable": true}, "action_parameters": {"description": "If step type is rule, then action parameters of this rule as JSON", "type": "object", "example": {"id": "4814384a9e5d4991b9815dcfc25d2f1f"}, "x-auditable": true}, "description": {"description": "If step type is rule or ruleset, the description of this entity", "type": "string", "example": "some rule", "x-auditable": true}, "expression": {"description": "If step type is rule, then expression used to match for this rule", "type": "string", "example": "ip.src ne 1.1.1.1", "x-auditable": true}, "kind": {"description": "If step type is ruleset, then kind of this ruleset", "type": "string", "example": "zone", "x-auditable": true}, "matched": {"description": "Whether tracing step affected tracing request/response", "type": "boolean", "example": true, "x-auditable": true}, "name": {"description": "If step type is ruleset, then name of this ruleset", "type": "string", "example": "some ruleset name", "x-auditable": true}, "step_name": {"description": "Tracing step identifying name", "type": "string", "example": "rule_id01", "x-auditable": true}, "trace": {"$ref": "#/components/schemas/request-tracer_trace"}, "type": {"description": "Tracing step type", "type": "string", "example": "rule", "x-auditable": true}}, "type": "object"}}
```
