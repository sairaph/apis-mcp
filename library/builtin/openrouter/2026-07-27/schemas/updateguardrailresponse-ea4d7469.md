---
title: UpdateGuardrailResponse
page_id: schema-updateguardrailresponse-ea4d7469
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UpdateGuardrailResponse

```yaml
{"example": {"data": {"allowed_models": null, "allowed_providers": ["openai"], "content_filter_builtins": [{"action": "redact", "label": "[EMAIL]", "slug": "email"}], "content_filters": null, "created_at": "2025-08-24T10:30:00Z", "description": "Updated description", "enforce_zdr": null, "enforce_zdr_anthropic": true, "enforce_zdr_google": true, "enforce_zdr_openai": true, "enforce_zdr_other": true, "enforce_zdr_xai": true, "id": "550e8400-e29b-41d4-a716-446655440000", "ignored_models": null, "ignored_providers": null, "include_byok_in_budgets": true, "limit_usd": 75, "name": "Updated Guardrail Name", "reset_interval": "weekly", "updated_at": "2025-08-24T16:00:00Z", "workspace_id": "0df9e665-d932-5740-b2c7-b52af166bc11"}}, "properties": {"data": {"allOf": [{"$ref": "#/components/schemas/Guardrail"}, {"description": "The updated guardrail"}]}}, "required": ["data"], "type": "object"}
```
