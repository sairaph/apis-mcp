---
title: NamespaceFunctionTool
page_id: schema-namespacefunctiontool-b0037d03
path: schemas
description: A function tool grouped inside a namespace tool
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# NamespaceFunctionTool

A function tool grouped inside a namespace tool

```yaml
{"description": "A function tool grouped inside a namespace tool", "example": {"name": "spawn_agent", "type": "function"}, "properties": {"allowed_callers": {"items": {"enum": ["direct", "programmatic"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": ["array", "null"]}, "defer_loading": {"type": "boolean"}, "description": {"type": ["string", "null"]}, "name": {"type": "string"}, "output_schema": {"additionalProperties": {}, "type": ["object", "null"]}, "parameters": {"additionalProperties": {}, "type": ["object", "null"]}, "strict": {"type": ["boolean", "null"]}, "type": {"enum": ["function"], "type": "string"}}, "required": ["type", "name"], "type": "object"}
```
