---
title: ListKeyAssignmentsResponse
page_id: schema-listkeyassignmentsresponse-fc97d305
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListKeyAssignmentsResponse

```yaml
{"example": {"data": [{"assigned_by": "user_abc123", "created_at": "2025-08-24T10:30:00Z", "guardrail_id": "550e8400-e29b-41d4-a716-446655440001", "id": "550e8400-e29b-41d4-a716-446655440000", "key_hash": "c56454edb818d6b14bc0d61c46025f1450b0f4012d12304ab40aacb519fcbc93", "key_label": "prod-key", "key_name": "Production Key"}], "total_count": 1}, "properties": {"data": {"description": "List of key assignments", "items": {"$ref": "#/components/schemas/KeyAssignment"}, "type": "array"}, "total_count": {"description": "Total number of key assignments for this guardrail", "example": 25, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
