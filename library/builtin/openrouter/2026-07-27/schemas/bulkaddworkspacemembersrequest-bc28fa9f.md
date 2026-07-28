---
title: BulkAddWorkspaceMembersRequest
page_id: schema-bulkaddworkspacemembersrequest-bc28fa9f
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# BulkAddWorkspaceMembersRequest

```yaml
{"example": {"user_ids": ["user_abc123", "user_def456"]}, "properties": {"user_ids": {"description": "List of user IDs to add to the workspace. Members are assigned the same role they hold in the organization.", "example": ["user_abc123", "user_def456"], "items": {"type": "string"}, "maxItems": 100, "minItems": 1, "type": "array"}}, "required": ["user_ids"], "type": "object"}
```
