---
title: ListMemberAssignmentsResponse
page_id: schema-listmemberassignmentsresponse-d126d3e7
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListMemberAssignmentsResponse

```yaml
{"example": {"data": [{"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "organization_id": "org_xyz789", "user_id": "user_abc123"}], "total_count": 1}, "properties": {"data": {"description": "List of member assignments", "items": {"$ref": "#/components/schemas/MemberAssignment"}, "type": "array"}, "total_count": {"description": "Total number of member assignments", "example": 10, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
