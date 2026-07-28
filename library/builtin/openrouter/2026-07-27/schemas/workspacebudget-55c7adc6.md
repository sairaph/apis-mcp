---
title: WorkspaceBudget
page_id: schema-workspacebudget-55c7adc6
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# WorkspaceBudget

```yaml
{"example": {"created_at": "2025-08-24T10:30:00Z", "id": "770e8400-e29b-41d4-a716-446655440000", "limit_usd": 100, "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}, "properties": {"created_at": {"description": "ISO 8601 timestamp of when the budget was created", "example": "2025-08-24T10:30:00Z", "type": "string"}, "id": {"description": "Unique identifier for the budget", "example": "770e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}, "limit_usd": {"description": "Spending limit in USD for this interval", "example": 100, "format": "double", "type": "number"}, "reset_interval": {"description": "Interval at which spend resets. Null means a lifetime (one-time) budget.", "enum": ["daily", "weekly", "monthly", null], "example": "monthly", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "updated_at": {"description": "ISO 8601 timestamp of when the budget was last updated", "example": "2025-08-24T15:45:00Z", "type": "string"}, "workspace_id": {"description": "ID of the workspace the budget belongs to", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, "required": ["id", "workspace_id", "limit_usd", "reset_interval", "created_at", "updated_at"], "type": "object"}
```
