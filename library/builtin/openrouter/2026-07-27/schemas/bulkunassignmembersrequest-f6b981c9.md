---
title: BulkUnassignMembersRequest
page_id: schema-bulkunassignmembersrequest-f6b981c9
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BulkUnassignMembersRequest

```yaml
{"example": {"member_user_ids": ["user_abc123", "user_def456"]}, "properties": {"member_user_ids": {"description": "Array of member user IDs to unassign from the guardrail", "example": ["user_abc123", "user_def456"], "items": {"minLength": 1, "type": "string"}, "minItems": 1, "type": "array"}}, "required": ["member_user_ids"], "type": "object"}
```
