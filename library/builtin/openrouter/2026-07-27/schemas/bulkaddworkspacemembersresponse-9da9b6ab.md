---
title: BulkAddWorkspaceMembersResponse
page_id: schema-bulkaddworkspacemembersresponse-9da9b6ab
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BulkAddWorkspaceMembersResponse

```yaml
{"example": {"added_count": 1, "data": [{"created_at": "2025-08-24T10:30:00Z", "id": "660e8400-e29b-41d4-a716-446655440000", "role": "member", "user_id": "user_abc123", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}]}, "properties": {"added_count": {"description": "Number of workspace memberships created or updated", "example": 2, "type": "integer"}, "data": {"description": "List of added workspace memberships", "items": {"$ref": "#/components/schemas/WorkspaceMember"}, "type": "array"}}, "required": ["data", "added_count"], "type": "object"}
```
