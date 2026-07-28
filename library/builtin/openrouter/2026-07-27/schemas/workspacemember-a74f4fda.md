---
title: WorkspaceMember
page_id: schema-workspacemember-a74f4fda
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WorkspaceMember

```yaml
{"example": {"created_at": "2025-08-24T10:30:00Z", "id": "660e8400-e29b-41d4-a716-446655440000", "role": "member", "user_id": "user_abc123", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}, "properties": {"created_at": {"description": "ISO 8601 timestamp of when the membership was created", "example": "2025-08-24T10:30:00Z", "type": "string"}, "id": {"description": "Unique identifier for the workspace membership", "example": "660e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}, "role": {"description": "Role of the member in the workspace", "enum": ["admin", "member"], "example": "member", "type": "string", "x-speakeasy-unknown-values": "allow"}, "user_id": {"description": "Clerk user ID of the member", "example": "user_abc123", "type": "string"}, "workspace_id": {"description": "ID of the workspace", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, "required": ["id", "workspace_id", "user_id", "role", "created_at"], "type": "object"}
```
