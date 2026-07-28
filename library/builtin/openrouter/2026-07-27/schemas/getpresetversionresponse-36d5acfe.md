---
title: GetPresetVersionResponse
page_id: schema-getpresetversionresponse-36d5acfe
path: schemas
description: A single version of a preset.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GetPresetVersionResponse

A single version of a preset.

```yaml
{"description": "A single version of a preset.", "example": {"data": {"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}}, "properties": {"data": {"$ref": "#/components/schemas/PresetDesignatedVersion"}}, "required": ["data"], "type": "object"}
```
