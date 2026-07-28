---
title: PresetWithDesignatedVersion
page_id: schema-presetwithdesignatedversion-1f43fb8d
path: schemas
description: A preset with its currently designated version.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PresetWithDesignatedVersion

A preset with its currently designated version.

```yaml
{"allOf": [{"$ref": "#/components/schemas/Preset"}, {"properties": {"designated_version": {"$ref": "#/components/schemas/PresetDesignatedVersion"}}, "required": ["designated_version"], "type": "object"}], "description": "A preset with its currently designated version.", "example": {"created_at": "2026-04-20T10:00:00Z", "creator_user_id": "user_2dHFtVWx2n56w6HkM0000000000", "description": null, "designated_version": {"config": {"model": "openai/gpt-4o", "temperature": 0.7}, "created_at": "2026-04-20T10:00:00Z", "creator_id": "user_2dHFtVWx2n56w6HkM0000000000", "id": "550e8400-e29b-41d4-a716-446655440000", "preset_id": "650e8400-e29b-41d4-a716-446655440001", "system_prompt": "You are a helpful assistant.", "updated_at": "2026-04-20T10:00:00Z", "version": 1}, "designated_version_id": "550e8400-e29b-41d4-a716-446655440000", "id": "650e8400-e29b-41d4-a716-446655440001", "name": "my-preset", "slug": "my-preset", "status": "active", "status_updated_at": null, "updated_at": "2026-04-20T10:00:00Z", "workspace_id": "750e8400-e29b-41d4-a716-446655440002"}}
```
