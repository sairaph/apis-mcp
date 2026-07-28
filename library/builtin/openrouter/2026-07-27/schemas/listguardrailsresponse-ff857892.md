---
title: ListGuardrailsResponse
page_id: schema-listguardrailsresponse-ff857892
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ListGuardrailsResponse

```yaml
{"example": {"data": [{"allowed_models": null, "allowed_providers": ["openai", "anthropic", "google"], "content_filter_builtins": [{"action": "redact", "label": "[EMAIL]", "slug": "email"}], "content_filters": null, "created_at": "2025-08-24T10:30:00Z", "description": "Guardrail for production environment", "enforce_zdr": null, "enforce_zdr_anthropic": true, "enforce_zdr_google": false, "enforce_zdr_openai": true, "enforce_zdr_other": false, "enforce_zdr_xai": false, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": false, "limit_usd": 100, "name": "Production Guardrail", "reset_interval": "monthly", "updated_at": "2025-08-24T15:45:00Z", "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}], "total_count": 1}, "properties": {"data": {"description": "List of guardrails", "items": {"$ref": "#/components/schemas/Guardrail"}, "type": "array"}, "total_count": {"description": "Total number of guardrails", "example": 25, "type": "integer"}}, "required": ["data", "total_count"], "type": "object"}
```
