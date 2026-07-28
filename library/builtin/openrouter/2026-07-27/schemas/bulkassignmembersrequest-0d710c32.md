---
title: BulkAssignMembersRequest
page_id: schema-bulkassignmembersrequest-0d710c32
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BulkAssignMembersRequest

```yaml
{"example": {"member_user_ids": ["user_abc123", "user_def456"]}, "properties": {"member_user_ids": {"description": "Array of member user IDs to assign to the guardrail", "example": ["user_abc123", "user_def456"], "items": {"minLength": 1, "type": "string"}, "minItems": 1, "type": "array"}}, "required": ["member_user_ids"], "type": "object"}
```
