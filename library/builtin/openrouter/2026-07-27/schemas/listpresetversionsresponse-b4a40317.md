---
title: ListPresetVersionsResponse
page_id: schema-listpresetversionsresponse-b4a40317
path: schemas
description: A paginated list of preset versions.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListPresetVersionsResponse

A paginated list of preset versions.

```yaml
{"description": "A paginated list of preset versions.", "example": {"data": [{"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}], "total_count": 1}, "properties": {"data": {"items": {"$ref": "#/components/schemas/PresetDesignatedVersion"}, "type": "array"}, "total_count": {"type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
