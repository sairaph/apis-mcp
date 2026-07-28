---
title: ListWorkspaceBudgetsResponse
page_id: schema-listworkspacebudgetsresponse-41080ee4
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListWorkspaceBudgetsResponse

```yaml
{"example": {"data": [{"created_at": "2025-08-24T10:30:00Z", "id": "770e8400-e29b-41d4-a716-446655440000", "limit_usd": 100, "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}], "include_byok_in_budgets": false}, "properties": {"data": {"description": "List of budgets configured for the workspace", "items": {"$ref": "#/components/schemas/WorkspaceBudget"}, "type": "array"}, "include_byok_in_budgets": {"description": "Whether BYOK spend is included in the workspace budgets", "example": false, "type": "boolean"}}, "required": ["data", "include_byok_in_budgets"], "type": "object"}
```
