---
title: ListWorkspaceMembersResponse
page_id: schema-listworkspacemembersresponse-f05c7dc8
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListWorkspaceMembersResponse

```yaml
{"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "id": "660e8400-e29b-41d4-a716-446655440000", "role": "member", "user_id": "user_abc123", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "total_count": 1}, "properties": {"data": {"description": "List of workspace members", "items": {"$ref": "#/components/schemas/WorkspaceMember"}, "type": "array"}, "total_count": {"description": "Total number of members in the workspace", "example": 5, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
