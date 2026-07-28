---
title: MessagesFallbackParam
page_id: schema-messagesfallbackparam-9c177f8f
path: schemas
description: Fallback model to try when the primary model fails or refuses. Only the `model` field is supported; per-attempt overrides are rejected.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesFallbackParam

Fallback model to try when the primary model fails or refuses. Only the `model` field is supported; per-attempt overrides are rejected.

```yaml
{"additionalProperties": {}, "description": "Fallback model to try when the primary model fails or refuses. Only the `model` field is supported; per-attempt overrides are rejected.", "example": {"model": "claude-opus-4-8"}, "properties": {"model": {"type": "string"}}, "required": ["model"], "type": "object"}
```
