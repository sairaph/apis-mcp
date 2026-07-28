---
title: PresetDesignatedVersion
page_id: schema-presetdesignatedversion-d66844b5
path: schemas
description: A specific version of a preset, containing config and optional system prompt.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PresetDesignatedVersion

A specific version of a preset, containing config and optional system prompt.

```yaml
{"description": "A specific version of a preset, containing config and optional system prompt.", "example": {"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}, "properties": {"config": {"additionalProperties": {}, "type": "object"}, "created_at": {"type": "string"}, "creator_id": {"type": "string"}, "id": {"type": "string"}, "preset_id": {"type": "string"}, "system_prompt": {"type": ["string", "null"]}, "updated_at": {"type": "string"}, "version": {"type": "integer"}}, "required": ["id", "preset_id", "creator_id", "version", "system_prompt", "config", "created_at", "updated_at"], "type": ["object", "null"]}
```
