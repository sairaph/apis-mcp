---
title: GetBYOKKeyResponse
page_id: schema-getbyokkeyresponse-2f8ad315
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# GetBYOKKeyResponse

```yaml
{"example": {"data": {"allowed_api_key_hashes": null, "allowed_models": null, "allowed_user_ids": null, "created_at": "2025-08-24T10:30:00Z", "disabled": false, "id": "11111111-2222-3333-4444-555555555555", "is_fallback": false, "label": "sk-...AbCd", "name": "Production OpenAI Key", "provider": "openai", "sort_order": 0, "workspace_id": "550e8400-e29b-41d4-a716-446655440000"}}, "properties": {"data": {"allOf": [{"$ref": "#/components/schemas/BYOKKey"}, {"description": "The BYOK credential."}]}}, "required": ["data"], "type": "object"}
```
