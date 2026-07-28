---
title: MemberAssignment
page_id: schema-memberassignment-03c19a91
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MemberAssignment

```yaml
{"example": {"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "organization_id": "org_xyz789", "user_id": "user_abc123"}, "properties": {"assigned_by": {"description": "User ID of who made the assignment", "example": "user_abc123", "type": ["string", "null"]}, "created_at": {"description": "ISO 8601 timestamp of when the assignment was created", "example": "2025-08-24T10:30:00Z", "type": "string"}, "guardrail_id": {"description": "ID of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440001", "format": "uuid", "type": "string"}, "id": {"description": "Unique identifier for the assignment", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}, "organization_id": {"description": "Organization ID", "example": "org_xyz789", "type": "string"}, "user_id": {"description": "Clerk user ID of the assigned member", "example": "user_abc123", "type": "string"}}, "required": ["id", "user_id", "organization_id", "guardrail_id", "assigned_by", "created_at"], "type": "object"}
```
