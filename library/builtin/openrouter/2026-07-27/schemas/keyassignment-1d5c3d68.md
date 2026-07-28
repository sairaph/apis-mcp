---
title: KeyAssignment
page_id: schema-keyassignment-1d5c3d68
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# KeyAssignment

```yaml
{"example": {"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "key_hash": "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93", "key_label": "prod-key", "key_name": "Production Key"}, "properties": {"assigned_by": {"description": "User ID of who made the assignment", "example": "user_abc123", "type": ["string", "null"]}, "created_at": {"description": "ISO 8601 timestamp of when the assignment was created", "example": "2025-08-24T10:30:00Z", "type": "string"}, "guardrail_id": {"description": "ID of the guardrail", "example": "550e8400-e29b-41d4-a716-446655440001", "format": "uuid", "type": "string"}, "id": {"description": "Unique identifier for the assignment", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}, "key_hash": {"description": "Hash of the assigned API key", "example": "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93", "type": "string"}, "key_label": {"description": "Label of the API key", "example": "prod-key", "type": "string"}, "key_name": {"description": "Name of the API key", "example": "Production Key", "type": "string"}}, "required": ["id", "key_hash", "guardrail_id", "key_name", "key_label", "assigned_by", "created_at"], "type": "object"}
```
