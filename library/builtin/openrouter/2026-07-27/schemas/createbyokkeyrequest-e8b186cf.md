---
title: CreateBYOKKeyRequest
page_id: schema-createbyokkeyrequest-e8b186cf
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CreateBYOKKeyRequest

```yaml
{"example": {"key": "sk-proj-abc123...", "name": "Production OpenAI Key", "provider": "openai"}, "properties": {"allowed_models": {"description": "Optional allowlist of model slugs this credential may be used for. `null` means no restriction.", "example": null, "items": {"type": "string"}, "maxItems": 100, "type": ["array", "null"]}, "allowed_user_ids": {"description": "Optional allowlist of user IDs that may use this credential. `null` means no restriction.", "example": null, "items": {"type": "string"}, "maxItems": 100, "type": ["array", "null"]}, "disabled": {"description": "Whether this credential should be created in a disabled state.", "example": false, "type": "boolean"}, "is_fallback": {"description": "Whether this credential is treated as a fallback — used only after non-fallback keys for the same provider have been tried.", "example": false, "type": "boolean"}, "key": {"description": "The raw provider API key or credential. This value is encrypted at rest and never returned in API responses.", "example": "sk-proj-abc123...", "minLength": 1, "type": "string"}, "name": {"description": "Optional human-readable name for the credential.", "example": "Production OpenAI Key", "maxLength": 255, "type": ["string", "null"]}, "provider": {"$ref": "#/components/schemas/BYOKProviderSlug"}, "workspace_id": {"description": "Optional workspace ID. Defaults to the authenticated entity's default workspace.", "example": "550e8400-e29b-41d4-a716-446655440000", "format": "uuid", "type": "string"}}, "required": ["provider", "key"], "type": "object"}
```
