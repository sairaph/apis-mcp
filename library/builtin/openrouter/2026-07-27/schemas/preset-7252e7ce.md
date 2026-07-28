---
title: Preset
page_id: schema-preset-7252e7ce
path: schemas
description: A preset without version details.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# Preset

A preset without version details.

```yaml
{"description": "A preset without version details.", "example": {"created_at": "2026-04-20T10:00:00Z", "creator_user_id": "user_2dHFtVWx2n56w6HkM0000000000", "description": null, "designated_version_id": "550e8400-e29b-41d4-a716-446655440000", "id": "650e8400-e29b-41d4-a716-446655440001", "name": "my-preset", "slug": "my-preset", "status": "active", "status_updated_at": null, "updated_at": "2026-04-20T10:00:00Z", "workspace_id": "750e8400-e29b-41d4-a716-446655440002"}, "properties": {"created_at": {"type": "string"}, "creator_user_id": {"type": ["string", "null"]}, "description": {"type": ["string", "null"]}, "designated_version_id": {"type": ["string", "null"]}, "id": {"type": "string"}, "name": {"type": "string"}, "slug": {"type": "string"}, "status": {"$ref": "#/components/schemas/PresetStatus"}, "status_updated_at": {"type": ["string", "null"]}, "updated_at": {"type": "string"}, "workspace_id": {"type": ["string", "null"]}}, "required": ["id", "creator_user_id", "workspace_id", "name", "slug", "description", "status", "designated_version_id", "created_at", "updated_at", "status_updated_at"], "type": "object"}
```
