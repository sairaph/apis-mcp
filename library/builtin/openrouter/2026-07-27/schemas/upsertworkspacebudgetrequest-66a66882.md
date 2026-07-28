---
title: UpsertWorkspaceBudgetRequest
page_id: schema-upsertworkspacebudgetrequest-66a66882
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpsertWorkspaceBudgetRequest

```yaml
{"example": {"include_byok_in_budgets": true, "limit_usd": 100}, "properties": {"include_byok_in_budgets": {"description": "Whether to include BYOK spend in the workspace budget", "example": true, "type": "boolean"}, "limit_usd": {"description": "Spending limit in USD. Must be greater than 0.", "example": 100, "format": "double", "type": "number"}}, "required": ["limit_usd"], "type": "object"}
```
