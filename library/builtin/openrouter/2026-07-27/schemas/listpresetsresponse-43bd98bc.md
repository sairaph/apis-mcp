---
title: ListPresetsResponse
page_id: schema-listpresetsresponse-43bd98bc
path: schemas
description: A paginated list of presets.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListPresetsResponse

A paginated list of presets.

```yaml
{"description": "A paginated list of presets.", "example": {"data": [{"created_at": "2026-04-20T10:00:00Z", "creator_user_id": "user_2dHFtVWx2n56w6HkM0000000000", "description": null, "designated_version_id": "550e8400-e29b-41d4-a716-446655440000", "id": "650e8400-e29b-41d4-a716-446655440001", "name": "my-preset", "slug": "my-preset", "status": "active", "status_updated_at": null, "updated_at": "2026-04-20T10:00:00Z", "workspace_id": "750e8400-e29b-41d4-a716-446655440002"}], "total_count": 1}, "properties": {"data": {"items": {"$ref": "#/components/schemas/Preset"}, "type": "array"}, "total_count": {"type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
